package Service

import (
	"T4/model"
	"T4/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// @title 发出留言
// @description 获取留言内容传给下一级发出留言函数
// @param c *gin.Context "对应的上下文"
// @return code int "状态码"
func PostMessage(c *gin.Context) (code int) {
	var msg model.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		log.Println(err)
		return emsg.Error
	}
	return model.PostMessage(&msg)
}

// @title 获取留言对象
// @description 获取留言id传给下一级获取留言函数
// @param c *gin.Context "对应的上下文"
// @return msg *model.Message "留言id对应的留言内容" code int "状态码"
func GetMessage(c *gin.Context) (*model.Message, int) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	return model.GetMessage("id", id)
}

// @title 删除留言
// @description 获取留言id传给下一级删除留言函数
// @param c *gin.Context "对应的上下文"
// @return code int "状态码"
func DeleteMessage(c *gin.Context) (code int) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	return model.DeleteMessage(id)
}

// @title 获取留言对象
// @description 获取回复的留言的id传给下一级获取留言函数
// @param c *gin.Context "对应的上下文"
// @return msg *model.Message "回复的留言的id对应的回复内容" code int "状态码"
func GetReply(c *gin.Context) (*model.Message, int) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return nil, emsg.Error
	}
	return model.GetMessage("reply_id", id)
}
