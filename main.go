package main

import (
	"github.com/globalsign/mgo"
	//"github.com/globalsign/mgo/bson"
	//"fmt"
	"App"
	"gee"
	"net/http"
	
)


const (
	url string = "127.0.0.1:27017" //mongo数据库连接端口
)


func initDB() {
	session, err := mgo.Dial(url) 
	if (err != nil) {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	App.MyuserModel = &App.UserModel {
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

	user := r.Group("/user")
	{
		user.POST("/register", App.RegisterUser)

		user.POST("/login",  App.LoginUser)

		user.GET("/:uid", App.GetUserByUid)

		user.POST("/:uid", App.ModifyUserByUid)

		user.GET("/:username", App.GetUserByUsername)
	}
	
	r.Run(":9999")
}