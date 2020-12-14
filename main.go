package main

import (
	"github.com/globalsign/mgo"
	//"github.com/globalsign/mgo/bson"
	//"fmt"
	"App"
	//"net/http"
	"github.com/gin-gonic/gin"
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
	App.MyarticleModel = &App.ArticleModel {
		DB: session.DB("myblog").C("article"),
	}
}

func main() {
	//连接数据库
	initDB()

	//开启服务器
	r := gin.Default()
	
	
	user := r.Group("/user")
	{
		user.POST("/register", App.RegisterUser)

		user.POST("/login",  App.LoginUser)

		user.GET("/username/:username", App.GetUserByUsername)

		user.GET("/uid/:uid", App.GetUserByUid)

		user.PUT("/:uid", App.ModifyUserByUid)

		
	}
	// router.Run()
	// // Simple group: v2
	// v2 := router.Group("/v2")
	// {
	// 	v2.POST("/login", loginEndpoint)
	// 	v2.POST("/submit", submitEndpoint)
	// 	v2.POST("/read", readEndpoint)
	// }

	
	// r := gee.New()
	// r.GET("/", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Welcome to the main page of myBlog!!!</h1>")
	// })

	// user := r.Group("/user")
	// {
	// 	user.POST("/register", App.RegisterUser)

	// 	user.POST("/login",  App.LoginUser)

	// 	user.GET("/:uid", App.GetUserByUid)

	// 	user.POST("/:uid", App.ModifyUserByUid)

	// 	user.GET("/username/:username", App.GetUserByUsername)
	// }

	// article := r.Group("/article")
	// {
	// 	article.POST("/publish", App.PublishArticle)

	// 	article.GET("/delete/:aid", App.DeleteArticleByAid)

	// 	article.POST("/:aid", App.ModifyArticleByAid)

	// 	article.GET("/all", App.GetAllArticles)

	// 	article.GET("/:aid", App.GetArticleByAid)

	// 	//article.GET("/title/:title", App.GetArticlesByTitle)

	// 	article.GET("/tag/:tag", App.GetArticlesByTag)

	// 	// article.GET("/publisher/:publisher", App.GetArticlesByPublisher)

	// }
	
	r.Run(":9999")
}