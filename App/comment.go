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

type CommentModel struct {
	DB *mgo.Collection
}

type Comment struct {
	Cid string   `bson:"_id"`
	Content string `bson:"content"`//评论内容
	Publisher string `bson:"publisher"`//评论者用户名
	Id string `bson:"id"` //评论的（文章或评论）的ID
}


//评论发布前没有Id
type Comment_notPublished struct {
	Content string `bson:"content"`//评论内容
	Publisher string `bson:"publisher"`//评论者用户名
	Id string `bson:"id"` //评论的（文章或评论）的ID
}

var (
	MycommentModel *CommentModel
)

//根据文章或评论ID查询评论
func GetCommentsById (c *gin.Context) {
	var comments []Comment
	err := MycommentModel.DB.Find(bson.M{"id": c.Param("id")}).All(&comments)
	if (err != nil) {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "",
		})
	} else {
		for i,comment := range comments {
			comments[i].Cid = fmt.Sprintf("%x", string(comment.Cid))
		}
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  comments,
		})
	}
}

//发表一条评论
func AddComment(c *gin.Context) {
	var comment Comment_notPublished
	json.NewDecoder(c.Request.Body).Decode(&comment)
	if comment.Id == "" {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "badparam",
		})
		return
	}
	var articles []Article
	var comments []Comment
	MyarticleModel.DB.Find(bson.M{"_id": bson.ObjectIdHex(comment.Id)}).All(&articles)
	MycommentModel.DB.Find(bson.M{"_id": bson.ObjectIdHex(comment.Id)}).All(&comments)
	if len(articles) == 0 && len(comments) == 0 {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "The contentID does not exist",
		})
	}else {
		MycommentModel.DB.Insert(&comment)
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  "",
		})
	}
}

//删除一条评论
func DeleteCommentByCid (c *gin.Context) {
	err := MycommentModel.DB.Remove(bson.M{"_id": bson.ObjectIdHex(c.Param("id"))})
	if err != nil {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "",
		})
	} else {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  "",
		})
	}


}

//修改一条评论
func ModifyCommentByCid (c *gin.Context) {
	var newcomment Comment
	json.NewDecoder(c.Request.Body).Decode(&newcomment)
	MycommentModel.DB.Update(bson.M{"_id": bson.ObjectIdHex(c.Param("id"))}, bson.M{"$set": bson.M{
		"content": newcomment.Content,
	}})
	c.JSON(http.StatusOK, &ApiResponse {
		Code: 200,
		Type: "success",
		Message:  "",
	})
}