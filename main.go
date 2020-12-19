package main

import (
	"github.com/globalsign/mgo"
	//"github.com/globalsign/mgo/bson"
	//"fmt"
	"App"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)


const (
	url string = "127.0.0.1:27017" //mongo数据库连接端口
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
	r.Use(passjs())
	store := cookie.NewStore([]byte("secret"))    //change
	r.Use(sessions.Sessions("sessionid",store))   //change
	user := r.Group("/user")
	{
		user.POST("/register", App.RegisterUser)

		user.POST("/login",  App.LoginUser)

		user.GET("/username/:username", App.GetUserByUsername)

		user.GET("/uid/:uid", App.GetUserByUid)

		user.Use(Authorize())

		user.PUT("/:uid", App.ModifyUserByUid)

	}

	article := r.Group("/article")
	{
		article.GET("/all", App.GetAllArticles)

		article.GET("/aid/:aid", App.GetArticleByAid)

		article.GET("/title/:title", App.GetArticlesByTitle)

		article.GET("/tag/:tag", App.GetArticlesByTag)

		article.GET("/publisher/:publisher", App.GetArticlesByPublisher)

		article.Use(Authorize())

		article.POST("/publish", App.PublishArticle)

		article.DELETE("/:aid", App.DeleteArticleByAid)

		article.PUT("/:aid", App.ModifyArticleByAid)

	}

	comment := r.Group("/comment")
	{
		comment.GET("/id/:id", App.GetCommentsById)

		comment.Use(Authorize())

		comment.POST("/publish", App.AddComment)

		comment.PUT("/:id", App.ModifyCommentByCid)

		comment.DELETE("/:id", App.DeleteCommentByCid)

	}

	like := r.Group("/like")
	{
		like.Use(Authorize())

		like.POST("/likeit", App.LikeIt)

		like.DELETE("/:lid", App.UnlikeIt)

		like.GET("/id/:id", App.GetLikesById)
	}
	
	r.Run(":9999")
}

func Authorize() gin.HandlerFunc{
	return func(c *gin.Context){
		session := sessions.Default(c)
		v := session.Get("sessionid")
		
		if v != nil {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"访问未授权"})
			return
		}
	}
}

// 中间件,主要处理js访问时的跨域问题
func passjs() gin.HandlerFunc {
	return func(c *gin.Context) {
		// gin设置响应头，设置跨域
	   c.Header("Access-Control-Allow-Origin", "*")
	   c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	   c.Header("Access-Control-Allow-Headers", "Action, Module, X-PINGOTHER, Content-Type, Content-Disposition")
	   if (c.Request.Method == "OPTIONS") {
		   c.JSON(200, &App.ApiResponse {
			   Code: 200,
		   })
		   return
	   }
	   // c.Next()后就执行真实的路由函数
		c.Next()
	}
}