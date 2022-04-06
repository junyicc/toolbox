package db

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	ConnectTimeout = 5 * time.Second
	PingTimeout    = 5 * time.Second
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
	CloseTimeout   = 5 * time.Second

	MaxMongoPoolSize = 200
)

func InitMongoDB(uri string) (*mongo.Database, error) {
	c, err := connect(uri)
	if err != nil {
		return nil, err
	}
	fields := strings.Split(uri, "/")
	dbName := fields[len(fields)-1]
	return c.Database(dbName), nil
}

func connect(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()
	// connect
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(MaxMongoPoolSize))
	if err != nil {
		return nil, err
	}

	// ping
	ctx, cancel = context.WithTimeout(context.Background(), PingTimeout)
	defer cancel()
	err = c.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return c, nil
}
