package App

import (
	//"errors"
	"github.com/globalsign/mgo"
	//"github.com/globalsign/mgo/bson"
	//"encoding/json"
	//"io/ioutil"
	//"fmt"
	//"net/http"
	"github.com/gin-gonic/gin"
)

type CommentModel struct {
	DB *mgo.Collection
}

type Comment struct {
	Cid string   `bson:"_id"`
	Message string `bson:"content"`//评论内容
	Publisher string `bson:"publisher"`//评论者用户名
	ContentId string `bson:"id"` //评论的（文章或评论）的ID
}


//评论发布前没有Id
type Comment_notPublished struct {
	Message string `bson:"content"`//评论内容
	Publisher string `bson:"publisher"`//评论者用户名
	ContentId string `bson:"id"` //评论的（文章或评论）的ID
}

var (
	MycommentModel *CommentModel
)

//根据文章或评论ID查询评论
func GetCommentsById (c *gin.Context) {

}

//发表一条评论
func AddComment(c *gin.Context) {

}

//删除一条评论
func DeleteCommentByCid (c *gin.Context) {

}

//修改一条评论
func ModifyCommentByCid (c *gin.Context) {

}