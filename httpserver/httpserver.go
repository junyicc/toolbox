package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"toolbox/db"
	"toolbox/logging"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HttpServer struct {
	host       string
	port       string
	server     *http.Server
	onShutdown []func()
	db         *gorm.DB
}

func New(serverConfig *ServerConfig, handler http.Handler) *HttpServer {
	return &HttpServer{
		host: serverConfig.Host,
		port: serverConfig.Port,
		server: &http.Server{
			Addr:         serverConfig.Host + ":" + serverConfig.Port,
			Handler:      handler,
			ReadTimeout:  serverConfig.ReadTimeout,
			WriteTimeout: serverConfig.WriteTimeout,
			IdleTimeout:  serverConfig.IdleTimeout,
		},
	}
}

func (s *HttpServer) ListenAndServe() error {
	for _, shutdown := range s.onShutdown {
		s.server.RegisterOnShutdown(shutdown)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		rec, ok := <-sig
		if ok {
			logging.Info("receive signal, try shutdown gracefully", zap.String("signal", rec.String()))
			if err := s.Shutdown(); err != nil {
				logging.Fatal("server shutdown", zap.Error(err))
			}
		}
	}()
	logging.Info("server start...", zap.String("host", s.host), zap.String("port", s.port))
	return s.server.ListenAndServe()
}

func (s *HttpServer) RegisterOnShutdown(f func()) {
	s.onShutdown = append(s.onShutdown, f)
}

func (s *HttpServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	<-ctx.Done()
	logging.Info("server shutdown")
	return nil
}

func (s *HttpServer) InitDB(mysqlConfig *db.MysqlConfig) {
	s.db = db.GormMysql(mysqlConfig)
	if s.db == nil {
		logging.Fatal("fail to init db")
	}
	// close db when shutdown server
	s.RegisterOnShutdown(func() {
		if s.db == nil {
			return
		}
		sqlDB, _ := s.db.DB()
		sqlDB.Close()
	})
}

func (s *HttpServer) DB() *gorm.DB {
	return s.db
}
