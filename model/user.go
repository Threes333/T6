package model

import (
	"T4/util"
)

// User 用户对象，定义了用户的基础信息
type User struct {
	Id           int    `json:"id" gorm:"type:int; primary_key auto_increment"`
	UserName     string `json:"username" gorm:"type:varchar(25)"`
	PassWord     string `json:"password" gorm:"type:varchar(25)"`
	SecretAnswer string `json:"secret_answer" gorm:"type:varchar(25)"`
}

// @title 用户注册
// @description 接收用户名和密码进行有效验证并在数据库进行存储
// @param username,password string "要进行注册的用户名与密码"
// @return code int "状态码"
func Register(username, password, answer string) (code int) {
	var id int
	db.Model(User{}).Select("id").Where("user_name = ?", username).First(&id)
	if id == 0 {
		//该用户名可用
		user := &User{
			UserName:     username,
			PassWord:     password,
			SecretAnswer: answer,
		}
		db.Create(user)
		return emsg.Success
	}
	return emsg.UsernameExist
}

// @title 用户登录
// @description 接收用户名和密码进行有效验证
// @param username,password string "要进行登录的用户名与密码"
// @return code int "状态码"
func Login(username, password string) (code int) {
	var pw string
	db.Model(User{}).Select("pass_word").Where("user_name = ?", username).First(&pw)
	//判断用户是否存在
	if pw == "" {
		return emsg.UsernameNoExist
	} else if pw == password { //判断密码是否正确
		return emsg.Success
	}
	return emsg.PasswordWrong
}

// @title 重置密码
// @description 接收用户名,密码和密保答案进行有效验证
// @param username,password string "要进行登录的用户名与密码"
// @return code int "状态码"
func ResetPassword(username, newPassword, answer string) (code int) {
	var secretAnswer string
	db.Model(User{}).Select("secret_answer").Where("user_name = ?", username).First(&secretAnswer)
	//判断用户是否存在
	if secretAnswer == "" {
		return emsg.UsernameNoExist
	} else if secretAnswer == answer { //判断密保答案是否正确
		db.Model(User{}).Where("user_name = ?", username).Update("pass_word", newPassword)
		return emsg.Success
	}
	return emsg.SecretAnswerWrong
}
