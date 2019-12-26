package message

import (
	"HelloWorld/io/db"
	"HelloWorld/io/network/connect"
	"HelloWorld/io/network/route"
	"HelloWorldServer/model"
	"HelloWorldServer/model/online"
	"HelloWorldServer/packet/trait"
	"HelloWorldServer/packet/user/contacts/message"
	"time"
)

func init() {
	route.Register(message.SendTextMessage{}, actionSendTextMessage)
}

func actionSendTextMessage(sendTextMessage message.SendTextMessage, conn connect.Connector) {

	selfUser, _ := online.Auth(conn.GetId())
	recvUserId := sendTextMessage.ContactId

	uc := db.Model(&model.UserContact{}).Where("user_id", selfUser.Id).AndWhere("contact_id", recvUserId).Exists()
	if !uc {
		conn.Send(message.SendMessageFail{Fail: trait.Fail{Message: "找不到该好友"}})
		return
	}
	cu := db.Model(&model.UserContact{}).Where("user_id", recvUserId).AndWhere("contact_id", selfUser.Id).Exists()
	if !cu {
		conn.Send(message.SendMessageFail{Fail: trait.Fail{Message: "您不是对方的好友"}})
		return
	}
	// 先发给自己发送成功提示
	conn.Send(message.SendMessageSuccess{SendTextMessage: sendTextMessage, Id: 0, CreateTime: uint64(time.Now().Unix())})
	//发送给好友
	online.UserSendMessage(recvUserId, message.RecvTextMessage{CreateTime: uint64(time.Now().Unix()), Id: 0, ContactId: selfUser.Id, Content: sendTextMessage.Content})

}
