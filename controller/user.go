package controller

import (
	"T4/service"
	"T4/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户注册的处理器函数
func Register(c *gin.Context) {
	code := Service.Register(c)
	//判断是否能正确注册
	if code == emsg.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}

//用户登录的处理器函数
func Login(c *gin.Context) {
	code := Service.Login(c)
	//判断是否能正确登录
	if code == emsg.Success {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}

//重置密码的处理器函数
func ResetPassword(c *gin.Context) {
	code := Service.ResetPassword(c)
	if code != emsg.Error {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}
