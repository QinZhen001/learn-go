package jwt_op

import "golang.org/x/oauth2/jwt"

const (
	TokenExpried     = "Token已过期"
	TokenNotValidYet = "Token不再有效"
	TokenMalformed   = "Token非法"
	TokenInvalid     = "Token无效"
)

type CustomClaims struct {
	jwt.StandardClaims
}
