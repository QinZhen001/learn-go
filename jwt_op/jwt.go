package jwt_op

import (
	"learngo/internal"
	"learngo/log"

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

// 生成token
func (j *JWT) GenerateJWT(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(j.SiginKey)
	if err != nil {
		log.Logger.Error("生成JWT错误:" + err.Error())
		return "", err
	}
	return tokenStr, nil
}

// 解析token
func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {

}
