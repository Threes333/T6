package model

import (
	"T4/util"
	"gorm.io/gorm"
	"log"
)

// Message 留言对象，定义了留言的基础信息
type Message struct {
	Id          int    `json:"id" gorm:"type:int,primary key auto_increment"` //留言的id
	Recipient   string `json:"recipient" gorm:"type:varchar(25)"`             //接受者的名字
	RecipientId int    `json:"recipient_id" gorm:"type:int"`                  //接收者的id
	Sender      string `json:"sender" gorm:"type:varchar(25)"`                //留言者的名字
	SenderId    int    `json:"sender_id" gorm:"type:int"`                     //留言者的id
	ReplyId     int    `json:"reply_id" gorm:"type:int"`                      //回复的留言的id
	Content     string `json:"content" gorm:"type:varchar(200)"`              //留言内容
}

// @title 发出留言
// @description 接收留言对象在数据库进行存储
// @param msg *Message "要存储的留言对象"
// @return code int "状态码"
func PostMessage(msg *Message) (code int) {
	if err := db.Create(msg).Error; err != nil {
		log.Println(err)
		return emsg.PostMessageFailed
	}
	var recipientId int
	if err := db.Model(User{}).Where("id = ?", msg.RecipientId).First(&recipientId).Error; err == gorm.ErrRecordNotFound {
		return emsg.UsernameNoExist
	}
	return emsg.Success
}

// @title 获取留言对象
// @description 接收留言id在数据库中查找返回留言内容
// @param id int "要查询的留言id" idType string "若要查询的是留言id则为id,若为回复留言的id即为reply_id"
// @return msg *Message "留言id对应的留言内容" code int "状态码"
func GetMessage(idType string, id int) (*Message, int) {
	var (
		msg Message
		err error
	)
	if idType == "id" {
		err = db.Debug().Model(Message{}).Where("id = ?", id).First(&msg).Error
	} else {
		err = db.Model(Message{}).Where("reply_id = ?", id).First(&msg).Error
	}
	if err == gorm.ErrRecordNotFound {
		return nil, emsg.MessageNoExist
	} else if err != nil {
		return nil, emsg.GetMessageFailed
	}
	return &msg, emsg.Success
}

// @title 获取留言对象
// @description 接收留言的留言id在数据库中查找返回留言内容
// @param id int "要查询的留言的留言id"
// @return msg *Message "留言的留言id对应的留言内容" code int "状态码"
func GetMessageByToMessageId(id int) (*Message, int) {
	var msg Message
	if err := db.Where("to_message_id = ?", id).Preload("Category").First(&msg).Error; err == gorm.ErrRecordNotFound {
		return nil, emsg.MessageNoExist
	} else if err != nil {
		return nil, emsg.GetMessageFailed
	}
	return &msg, emsg.Success
}

// @title 删除留言
// @description 接收留言id将其删除
// @param id int "要删除的留言id"
// @return code int "状态码"
func DeleteMessage(id int) (code int) {
	if err := db.Where("id = ?", id).Delete(&Message{}).Error; err != nil {
		log.Println(err)
		return emsg.DeleteMessageFailed
	}
	return emsg.Success
}
