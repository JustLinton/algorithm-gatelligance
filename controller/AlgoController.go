package controller

import (
	"fmt"
	"net/http"

	"gatelligance_algo/service"

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

		tid := service.CreateLinkTransaction(db, addr, err)
		c.String(http.StatusOK, tid)
	})

	//for test
	router.GET("/sayHello", func(c *gin.Context) {

		// strToken := c.DefaultQuery("token", "nil")
		// claim, stat := Verification.VerifyToken(strToken)
		// if !stat {
		// 	c.String(http.StatusOK, "Login expired.")
		// 	return
		// }
		// c.String(http.StatusOK, "Hello,"+claim.ID)

	})
}
