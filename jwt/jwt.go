package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	UID      = "uid"
	Username = "username"
)

var (
	defaultJWTKey           string = `QlMTdST4vyVn5WyKdWBJkU9YHANMNb2yXinhIl6ZHgUh2GCcTLQCkRU2vmnfhW5`
	defaultJWTExpireSeconds int64  = 604800 // 7 days
)

// Claims 声明
type Claims struct {
	jwt.StandardClaims

	UID      uint   `json:"uid"`
	Username string `json:"username"`
}

func CreateClaims(uid uint, username string) Claims {
	return Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + defaultJWTExpireSeconds,
		},
		UID:      uid,
		Username: username,
	}
}

// GenToken generates jwt token
func GenToken(claims Claims, key []byte) (string, error) {
	if len(key) == 0 {
		key = []byte(defaultJWTKey)
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 获取完整签名之后的 token
	return tokenClaims.SignedString(key)
}

// ParseToken parse jwt token
func ParseToken(token string, key []byte) (*Claims, error) {
	if len(key) == 0 {
		key = []byte(defaultJWTKey)
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
	if err != nil {
		var jwtErr *jwt.ValidationError
		if errors.As(err, &jwtErr) {
			if jwtErr.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			}
			if jwtErr.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			}
			if jwtErr.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			}
			return nil, ErrTokenInvalid
		}
	}

	if jwtToken != nil {
		if claims, ok := jwtToken.Claims.(*Claims); ok && jwtToken.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid
	}
	return nil, ErrTokenInvalid
}
