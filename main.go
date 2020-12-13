package main

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"fmt"
	"App"
	"gee"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type ApiResponse struct {
	Code int  `bson:"code"`
	Type string  `bson:"type"`
	Message string  `bson:"message"`
}

const (
	url string = "127.0.0.1:27017" //mongo数据库连接端口
)

var (
	userModel *App.UserModel
)

func initDB() {
	session, err := mgo.Dial(url) 
	if (err != nil) {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	userModel = &App.UserModel {
		DB: session.DB("myblog").C("user"),
	}
}

func main() {
	//连接数据库
	initDB()
	

	//开启服务器
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Welcome to the main page of myBlog!!!</h1>")
	})

	v2 := r.Group("/user")
	{
		v2.POST("/register", func(c *gee.Context) {
			//解析post的数据存到postUser内
			con,_ := ioutil.ReadAll(c.Req.Body) //获取post的数据
			postUser := App.User_notRegistered{}
			json.Unmarshal(con, &postUser)

			//检查用户名是否已经被注册
			tmpUser := App.User{}
			userModel.DB.Find(bson.M{"username": postUser.Username}).One(&tmpUser)
			hexid := fmt.Sprintf("%x", string(tmpUser.Id))
			if (hexid == "") {
				err := userModel.DB.Insert(&postUser)
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
		})

		v2.POST("/login", func(c *gee.Context) {
			//解析post的数据存到postUser内
			con,_ := ioutil.ReadAll(c.Req.Body) //获取post的数据
			postUser := App.User{}
			json.Unmarshal(con, &postUser)

			//检查用户名和密码是否匹配
			tmpUser := App.User{}
			userModel.DB.Find(bson.M{"username": postUser.Username, "password": postUser.Password}).One(&tmpUser)
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
		})
	}
	
	r.Run(":9999")
}