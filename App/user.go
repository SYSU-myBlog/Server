package App

import (
	//"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
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
	Id string   `bson:"_id"`
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

// //非本人只能查看以下信息
// type OtherUser struct {
// 	Username string	  `bson:"username"`
// 	Email string	`bson:"email"`
// 	UserStatus UserStatus  `bson:"userStatus"`
// }

var (
	MyuserModel *UserModel
)

//用户注册
func RegisterUser(c *gin.Context) {
	//解析post的数据存到postUser内
	con,_ := ioutil.ReadAll(c.Request.Body) //获取post的数据
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
func LoginUser (c *gin.Context) {
	session := sessions.Default(c)                //change
	option := sessions.Options{MaxAge: 3600, Path: "/"}      //change
	session.Options(option)                       //change
	//解析post的数据存到postUser内
	con,_ := ioutil.ReadAll(c.Request.Body) //获取post的数据
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
		session.Set("sessionid", tmpUser.Id)       //change
		session.Save()                             //change
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  &ObjectID {
				Id: hexid,
			},
		})
	}
}

// GetUserByID 根据ID查询用户
func GetUserByUid (c *gin.Context) {
	tmpUser := User{}
	MyuserModel.DB.FindId(bson.ObjectIdHex(c.Param("uid"))).One(&tmpUser)
	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
	if (hexid == "") {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "user id does not exist",
		})
	} else {
		tmpUser.Id = fmt.Sprintf("%x", string(tmpUser.Id))
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  &tmpUser,
		})
	}
}

// 根据用户名查询用户
func GetUserByUsername (c *gin.Context) {
	tmpUser := User{}
	MyuserModel.DB.Find(bson.M{"username": c.Param("username")}).One(&tmpUser)
	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
	if (hexid == "") {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "user does not exist",
		})
	} else {
		tmpUser.Id = fmt.Sprintf("%x", string(tmpUser.Id))
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  &tmpUser,
		})
	}
}


//修改用户信息
func ModifyUserByUid (c *gin.Context) {
	//解析post的数据存到postUser内
	con,_ := ioutil.ReadAll(c.Request.Body) //获取post的数据
	postUser := User{}
	json.Unmarshal(con, &postUser)

	tmpUser := User{}
	MyuserModel.DB.FindId(bson.ObjectIdHex(c.Param("uid"))).One(&tmpUser)
	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
	if (hexid == "") {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "user does not exist",
		})
	} else {
		//更新
		if (postUser.Username == "") {
			postUser.Username = tmpUser.Username
		}
		if (postUser.Email == "") {
			postUser.Email = tmpUser.Email
		}
		if (postUser.Password == "") {
			postUser.Password = tmpUser.Password
		}
		if (postUser.Phone == "") {
			postUser.Phone = tmpUser.Phone
		}
		MyuserModel.DB.Update(bson.M{"_id": bson.ObjectIdHex(c.Param("uid"))}, bson.M{"$set": bson.M{
			"username": postUser.Username,
			"email": postUser.Email,
			"password": postUser.Password,
			"phone": postUser.Phone,
		}})
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  "modify user success",
		})
	}
	
}

