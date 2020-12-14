package App

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"gee"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
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


//未注册用户没有Id等其他属性
type User_notRegistered struct {
	Username string	  `bson:"username"`
	Password string		`bson:"password"`
}

var (
	MyuserModel *UserModel
)

//用户注册
func RegisterUser(c *gee.Context) {
	//解析post的数据存到postUser内
	con,_ := ioutil.ReadAll(c.Req.Body) //获取post的数据
	postUser := User_notRegistered{}
	json.Unmarshal(con, &postUser)

	//检查用户名是否已经被注册
	tmpUser := User{}
	MyuserModel.DB.Find(bson.M{"username": postUser.Username}).One(&tmpUser)
	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
	if (hexid == "") {
		err := MyuserModel.DB.Insert(&postUser)
		if (err != nil) {
			panic(err)
		}
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  "register success",
		})
	} else {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "username has existed.",
		})
	}
}

//用户登陆
func LoginUser (c *gee.Context) {
	//解析post的数据存到postUser内
	con,_ := ioutil.ReadAll(c.Req.Body) //获取post的数据
	postUser := User{}
	json.Unmarshal(con, &postUser)

	//检查用户名和密码是否匹配
	tmpUser := User{}
	MyuserModel.DB.Find(bson.M{"username": postUser.Username, "password": postUser.Password}).One(&tmpUser)
	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
	if (hexid == "") {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "username and password do not match",
		})
	} else {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  "login success.",
		})
	}
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

