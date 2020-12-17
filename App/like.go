package App

import (
	//"errors"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"encoding/json"
	//"io/ioutil"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type LikeModel struct {
	DB *mgo.Collection
}

type Like struct {
	Lid string   `bson:"_id"`
	Liker string `bson:"liker"`//点赞者用户名
	Id string `bson:"id"` //点赞的（文章或评论）的ID
}


//点赞发布前没有Id
type Like_notPublished struct {
	Liker string `bson:"liker"`//点赞者用户名
	Id string `bson:"id"` //点赞的（文章或评论）的ID
}

var (
	MylikeModel *LikeModel
)

//点赞
func LikeIt (c *gin.Context) {
	var like Like_notPublished
	json.NewDecoder(c.Request.Body).Decode(&like)
	var tmpUser User
	MyuserModel.DB.Find(bson.M{"username": like.Liker}).One(&tmpUser)
	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
	if (hexid == "") {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "user does not exist",
		})
		return 
	}
	if like.Id == "" {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "content does not exist",
		})
		return
	}
	var articles []Article
	var comments []Comment
	MyarticleModel.DB.Find(bson.M{"_id": bson.ObjectIdHex(like.Id)}).All(&articles)
	MycommentModel.DB.Find(bson.M{"_id": bson.ObjectIdHex(like.Id)}).All(&comments)
	if len(articles) == 0 && len(comments) == 0 {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "The contentID does not exist",
		})
	}else {
		MylikeModel.DB.Insert(&like)
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  "",
		})
	}
}

//取消点赞
func UnlikeIt (c *gin.Context) {
	var like []Like
	MylikeModel.DB.Find(bson.M{"_id": bson.ObjectIdHex(c.Param("lid"))}).All(&like)
	if len(like) == 0 {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "The contenID does not exist",
		})
	} else {
		MylikeModel.DB.Remove(bson.M{"_id": bson.ObjectIdHex(c.Param("lid"))})
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  "",
		})
	}
}

func GetLikesById (c *gin.Context) {
	var likes []Like
	err := MylikeModel.DB.Find(bson.M{"id": c.Param("id")}).All(&likes)
	if (err != nil) {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "",
		})
	} else {
		for i,like := range likes {
			likes[i].Lid = fmt.Sprintf("%x", string(like.Lid))
		}
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  likes,
		})
	}
}