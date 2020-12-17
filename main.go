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
	url string = "172.19.34.164:27017" //mongo数据库连接端口
	//url string = "172.26.43.243:27017" //mongo数据库连接端口
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
	App.MycommentModel  = &App.CommentModel {
		DB: session.DB("myblog").C("comment"),
	}
	App.MylikeModel  = &App.LikeModel {
		DB: session.DB("myblog").C("like"),
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

	article := r.Group("/article")
	{
		article.POST("/publish", App.PublishArticle)

		article.DELETE("/:aid", App.DeleteArticleByAid)

		article.PUT("/:aid", App.ModifyArticleByAid)

		article.GET("/all", App.GetAllArticles)

		article.GET("/aid/:aid", App.GetArticleByAid)

		article.GET("/title/:title", App.GetArticlesByTitle)

		article.GET("/tag/:tag", App.GetArticlesByTag)

		article.GET("/publisher/:publisher", App.GetArticlesByPublisher)

	}

	comment := r.Group("/comment")
	{
		comment.POST("/publish", App.AddComment)

		comment.GET("/id/:id", App.GetCommentsById)

		comment.PUT("/:id", App.ModifyCommentByCid)

		comment.DELETE("/:id", App.DeleteCommentByCid)

	}

	like := r.Group("/like")
	{
		like.POST("/likeit", App.LikeIt)

		like.DELETE("/:lid", App.UnlikeIt)

		like.GET("/id/:id", App.GetLikesById)
	}
	
	r.Run(":9999")
}