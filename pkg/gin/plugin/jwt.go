package plugin

import (
	"errors"
	"flyingv2/internal/core/resp"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/sync/singleflight"
	"strconv"
	"time"
)

const (
	req_un = "request unauthorized"
)

type Authentication struct {
	Secret     string
	ExpireTime int64 `yaml:"expire_time"`
}

var g singleflight.Group

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
	j                *Authentication
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" && len(token) <= 0 {
			resp.Unauthorized(req_un, c)
			c.Abort()
			return
		}

		if claims, err := j.ParseToken(token); err != nil {
			if err == TokenExpired {
				resp.Unauthorized(TokenExpired.Error(), c)
				c.Abort()
				return
			}
			resp.Unauthorized(err.Error(), c)
			c.Abort()
			return
		} else {
			if claims.ExpiresAt-time.Now().Unix() < 86400 {
				newToken, newClaims := j.RefreshToken(
					claims,
					token,
				)
				c.Header("new-token", newToken)
				c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			}
			c.Set("claims", claims)
			c.Next()
		}

	}

}
func init() {
	j = NewAuth()
}
func NewAuth() *Authentication {
	return &Authentication{
		Secret:     "ewfcdsceddsckxdfql",
		ExpireTime: 604800,
	}
}

// 创建一个token
func (j *Authentication) CreateToken(claims jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.Secret)
}

func (j *Authentication) ParseToken(tokenStr string) (*jwt.StandardClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}
}
func (j *Authentication) CreateTokenByOldToken(oldToken string, claims jwt.StandardClaims) (string, error) {
	v, err, _ := g.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}
func (j *Authentication) RefreshToken(claims *jwt.StandardClaims, tokenStr string) (string, *jwt.StandardClaims) {
	claims.ExpiresAt = time.Now().Unix() + j.ExpireTime
	newToken, _ := j.CreateTokenByOldToken(tokenStr, *claims)
	newClaims, _ := j.ParseToken(newToken)
	return newToken, newClaims
}
