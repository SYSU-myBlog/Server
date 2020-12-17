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
		//GetAllSonComments(c.Param("id"))
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  comments,
		})
	}
}

func GetCommentsByid (id string) []string {
	var comments []Comment
	var commentsid []string
	MycommentModel.DB.Find(bson.M{"id": id}).All(&comments)
	for i,comment := range comments {
		comments[i].Cid = fmt.Sprintf("%x", string(comment.Cid))
		commentsid = append(commentsid, comments[i].Cid)
	}
	return commentsid
}

func GetAllSonComments (id string) []string{
	var arrstr []string
	var res []string
	arrstr = append(arrstr, id)
	/*for _, v := range arrstr {
		arrstr = append(arrstr, GetCommentsByid(v)...)
		res = append(res, GetCommentsByid(v)...)
		arrstr = arrstr[1:] // 删除开头1个元素
		fmt.Println(len(arrstr))
	}*/
	i := 0
	res = append(res, id)
	for ; i < len(arrstr); i++ {
		res = append(res, GetCommentsByid(arrstr[i])...)
		arrstr = append(arrstr, GetCommentsByid(arrstr[i])...)	
		arrstr = arrstr[1:] // 删除开头1个元素
		i = -1
	}
	return res
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
	var tmpUser User
	MyuserModel.DB.Find(bson.M{"username": comment.Publisher}).One(&tmpUser)
	hexid := fmt.Sprintf("%x", string(tmpUser.Id))
	if (hexid == "") {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "publisher does not exist",
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
	var delecomments []string
	delecomments = append(delecomments,GetAllSonComments(c.Param("id"))...)
	for _,v := range delecomments {
		MycommentModel.DB.Remove(bson.M{"_id": bson.ObjectIdHex(v)})
		MylikeModel.DB.Remove(bson.M{"id": v})
	}
	c.JSON(http.StatusOK, &ApiResponse {
		Code: 200,
		Type: "success",
		Message:  "",
	})
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