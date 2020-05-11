package model

import "github.com/yanlong-li/HelloWorld-GO/io/db"

type UserContactBlack struct {
	Id         uint64
	UserId     uint64
	ContactId  uint64
	CreateTime uint64
	Remark     string
}

func (uc *UserContactBlack) GetUserInfo() (User, db.OrmError) {
	return GetUserById(uc.UserId)
}

func (uc *UserContactBlack) GetContactInfo() (User, db.OrmError) {
	return GetUserById(uc.ContactId)
}
