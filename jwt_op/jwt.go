package jwt_op

import (
	"learngo/internal"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenExpried     = "Token已过期"
	TokenNotValidYet = "Token不再有效"
	TokenMalformed   = "Token非法"
	TokenInvalid     = "Token无效"
)

type CustomClaims struct {
	jwt.StandardClaims
	Id          int32
	NikeName    string
	AuthorityId int32
}

type JWT struct {
	SiginKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		SiginKey: []byte(internal.AppConf.JWTConfig.SingingKey),
	}
}

func GenerateJWT() {}
