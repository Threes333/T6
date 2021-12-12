package controller

import (
	"T4/service"
	"T4/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

//发出留言的处理器函数
func PostMessage(c *gin.Context) {
	code := Service.PostMessage(c)
	if code != emsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}

//获取留言的处理器函数
func GetMessage(c *gin.Context) {
	data, code := Service.GetMessage(c)
	if code != emsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": *data,
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}

//删除留言的处理器函数
func DeleteMessage(c *gin.Context) {
	code := Service.DeleteMessage(c)
	if code != emsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}

//获取留言留言的处理器函数
func GetReply(c *gin.Context) {
	data, code := Service.GetReply(c)
	if code != emsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": *data,
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}
