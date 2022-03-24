package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	Service "gatelligance/service"
	Verification "gatelligance/verification"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitAlgoController(err *error, db *gorm.DB, router *gin.Engine) {

	//新增链接下载任务
	router.POST("/addLinkWork", func(c *gin.Context) {

		addr := c.DefaultPostForm("addr", "nil")
		id := c.DefaultPostForm("id", "nil")

		if addr == "nil" || id == "nil" {
			c.String(http.StatusNotAcceptable, fmt.Sprintln("network"))
			return
		}

		res := Service.GetAudioText(addr, id)
		c.String(http.StatusOK, res)
	})

	//for test
	router.GET("/frontEnd/sayHello", func(c *gin.Context) {

		strToken := c.DefaultQuery("token", "nil")
		claim, stat := Verification.VerifyToken(strToken)
		if !stat {
			c.String(http.StatusOK, "Login expired.")
			return
		}
		c.String(http.StatusOK, "Hello,"+claim.ID)
	})
}

func getSHA256HashCode(message []byte) string {
	hash := sha256.New()
	hash.Write(message)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
