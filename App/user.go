package App

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type UserModel struct {
	DB *mgo.Collection
}

type UserStatus int32
const (
	StatusOnline UserStatus = 0
	StatusOffline UserStatus = 1
)

type User struct {
	Id bson.ObjectId   `bson:"_id"`
	Username string	  `bson:"username"`
	Email string	`bson:"email"`
	Password string		`bson:"password"`
	Phone string	`bson:"phone"`
	UserStatus UserStatus  `bson:"userStatus"`
}

//未注册用户没有Id
type User_notRegistered struct {
	Username string	  `bson:"username"`
	Email string	`bson:"email"`
	Password string		`bson:"password"`
	Phone string	`bson:"phone"`
	UserStatus UserStatus  `bson:"userStatus"`
}

// GetUsers 获取所有用户
func (m *UserModel) GetUsers() (users []User, err error) {
	err = m.DB.Find(nil).All(&users)
	return
}

// GetUserByID 根据ID查询用户
func (m *UserModel) GetUserByID(id string) (user User, err error) {
	if !bson.IsObjectIdHex(id) {
		err = errors.New("not_id")
		return
	}
	err = m.DB.FindId(bson.ObjectIdHex(id)).One(&user)
	return
}

