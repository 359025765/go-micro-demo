package service

import (
	pb "github.com/359025765/go-micro-demo/user-service/proto/user"
	"github.com/359025765/go-micro-demo/user-service/repo"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	key = []byte("laracomUserTokenKeySecret")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type TokenService struct {
	Repo repo.Repository
}

// Decode a token string into a token object
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into a JWT
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "laracom.user.service",
		},
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名算法token并返回
	return token.SignedString(key)
}
