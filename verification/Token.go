package verification

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "登录已过期"
)

// func SayHello(c *gin.Context) {
// 	strToken := c.DefaultQuery("token", "nil")
// 	claim, err := VerifyAction(strToken)
// 	if err != nil {
// 		c.String(http.StatusNotFound, err.Error())
// 		return
// 	}
// 	c.String(http.StatusOK, "hello,", claim.ID)
// }

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	ID string `json:"id"`
}

var (
	Secret     = "this-is-a-secret-key-by-linton-jiang." // 加盐
	ExpireTime = 3600                                    // token有效期
)

func VerifyToken(strToken string) (*JWTClaims, bool) {
	claim, err := verifyAction(strToken)
	if err != nil {
		return claim, false
	}
	return claim, true
}

func VerifyTokenHandler(c *gin.Context) {
	strToken := c.Param("token")
	claim, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "verify,", claim.ID)
}

func RefreshTokenHandler(c *gin.Context) {

	strToken := c.DefaultQuery("token", "nil")
	claims, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := GetToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)

}

func verifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorReason_ServerBusy)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	fmt.Println("verify")
	return claims, nil
}

func GetToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	return signedToken, nil
}