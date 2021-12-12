package cmd

import (
	"T4/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	//用户相关
	r.POST("/register", controller.Register)
	r.GET("/login", controller.Login)
	r.PUT("/password", controller.ResetPassword)
	//留言相关
	r.GET("/message/:id", controller.GetMessage)
	r.POST("/message", controller.PostMessage)
	r.DELETE("/message/:id", controller.DeleteMessage)
	r.GET("/message/reply/:id", controller.GetReply)
	_ = r.Run(":8080")
}
