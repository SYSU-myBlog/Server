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

type LikeModel struct {
	DB *mgo.Collection
}

type Like struct {
	Lid string   `bson:"_id"`
	Liker string `bson:"liker"`//点赞者用户名
	ContentId string `bson:"id"` //点赞的（文章或评论）的ID
}


//点赞发布前没有Id
type Like_notPublished struct {
	Liker string `bson:"liker"`//点赞者用户名
	ContentId string `bson:"id"` //点赞的（文章或评论）的ID
}

var (
	MylikeModel *LikeModel
)

//点赞
func LikeIt (c *gin.Context) {

}

//取消点赞
func UnlikeIt (c *gin.Context) {

}