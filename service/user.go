package Service

import (
	"T4/model"
	"T4/util"
	"github.com/gin-gonic/gin"
	"log"
)

// @title 用户注册
// @description 获取用户名,密码和密保答案传给下一级用户注册函数
// @param c *gin.Context "相应的上下文"
// @return code int "状态码"
func Register(c *gin.Context) (code int) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	return model.Register(user.UserName, user.PassWord, user.SecretAnswer)
}

// @title 用户登录
// @description 获取用户名和密码传给下一级用户登录函数
// @param c *gin.Context "相应的上下文"
// @return code int "状态码"
func Login(c *gin.Context) (code int) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return emsg.Error
	}
	return model.Login(user.UserName, user.PassWord)
}

// @title 重置密码
// @description 获取用户名,新密码和密保答案传给下一级重置密码函数
// @param c *gin.Context "相应的上下文"
// @return code int "状态码"
func ResetPassword(c *gin.Context) (code int) {
	username := c.PostForm("username")
	newPassword := c.PostForm("new_password")
	secretAnswer := c.PostForm("secret_answer")
	if username == "" || newPassword == "" || secretAnswer == "" {
		return emsg.Error
	}
	return model.ResetPassword(username, newPassword, secretAnswer)
}
