package App

import (
	"github.com/globalsign/mgo"
	// "github.com/globalsign/mgo/bson"
	// "encoding/json"
	// "io/ioutil"
	// "fmt"
	// "net/http"
	// "time"
)

type Article struct {
	Aid string `bson:"_id"`
	Title string `bson:"title"`
	Content string `bson:"content"`
	Publisher string`bson:"publisher"`
	Tag string `bson:"tag"`
	LastModified string `bson:"lastModified"`
}

//文章发布前没有ID
type Article_notPublished struct {
	Title string `bson:"title"`
	Content string `bson:"content"`
	Publisher string `bson:"publisher"`
	Tag string `bson:"tag"`
	LastModified string `bson:"lastModified"`
}

type ArticleModel struct {
	DB *mgo.Collection
}

var (
	MyarticleModel *ArticleModel
)
/*
//获取所有文章
func GetAllArticles(c *gee.Context) {
	articles := []Article{}
	MyarticleModel.DB.Find(nil).All(&articles)
	c.JSON(http.StatusOK, &ApiResponse {
		Code: 200,
		Type: "success",
		Message:  &articles,
	})
}

//发表一篇文章
func PublishArticle(c *gee.Context) {
	//解析post的数据存到postArticle内
	con,_ := ioutil.ReadAll(c.Req.Body) //获取post的数据
	postArticle := Article_notPublished{}
	json.Unmarshal(con, &postArticle)

	postArticle.LastModified = time.Now().Format("2006-01-02 15:04:05")
	//插入数据
	MyarticleModel.DB.Insert(&postArticle)
	//查找插入的数据
	tmpArticle := Article{}
	MyarticleModel.DB.Find(bson.M{"title": postArticle.Title, "content": postArticle.Content}).One(&tmpArticle)
	hexid := fmt.Sprintf("%x", string(tmpArticle.Aid))

	c.JSON(http.StatusOK, &ApiResponse {
		Code: 200,
		Type: "success",
		Message:  &ObjectID{Id: hexid},
	})
}

//根据Id查找一篇文章
func GetArticleByAid(c *gee.Context) {
	tmpArticle := Article{}
	MyarticleModel.DB.FindId(bson.ObjectIdHex(c.Param("aid"))).One(&tmpArticle)
	hexid := fmt.Sprintf("%x", string(tmpArticle.Aid))
	if (hexid == "") {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "article id does not exist",
		})
	} else {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  &tmpArticle,
		})
	}
}

// 根据标题查找文章
func GetArticlesByTitle (c *gee.Context) {
	articles := []Article{}
	MyarticleModel.DB.Find(bson.M{"title": c.Param("title")}).All(&articles)
	c.JSON(http.StatusOK, &ApiResponse {
		Code: 200,
		Type: "success",
		Message:  &articles,
	})
}

// 根据标签查找文章
func GetArticlesByTag (c *gee.Context) {
	articles := []Article{}
	MyarticleModel.DB.Find(bson.M{"tag": c.Param("tag")}).All(&articles)
	c.JSON(http.StatusOK, &ApiResponse {
		Code: 200,
		Type: "success",
		Message:  &articles,
	})
}

// 根据用户名查找文章
func GetArticlesByPublisher (c *gee.Context) {
	articles := []Article{}
	MyarticleModel.DB.Find(bson.M{"publisher": c.Param("publisher")}).All(&articles)
	c.JSON(http.StatusOK, &ApiResponse {
		Code: 200,
		Type: "success",
		Message:  &articles,
	})
}



//根据Id删除一篇文章
func DeleteArticleByAid (c *gee.Context) {
	tmpArticle := Article{}
	MyarticleModel.DB.FindId(bson.ObjectIdHex(c.Param("aid"))).One(&tmpArticle)
	hexid := fmt.Sprintf("%x", string(tmpArticle.Aid))
	if (hexid == "") {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "article id does not exist",
		})
	} else {
		MyarticleModel.DB.Remove(bson.M{"_id": bson.ObjectIdHex(c.Param("aid"))})
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  "delete article success",
		})
	}
}

//根据Id修改一篇文章
func ModifyArticleByAid (c *gee.Context) {
	//解析post的数据存到postUser内
	con,_ := ioutil.ReadAll(c.Req.Body) //获取post的数据
	postArticle := Article{}
	json.Unmarshal(con, &postArticle)

	tmpArticle := Article{}
	MyarticleModel.DB.FindId(bson.ObjectIdHex(c.Param("aid"))).One(&tmpArticle)
	hexid := fmt.Sprintf("%x", string(tmpArticle.Aid))
	if (hexid == "") {
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 400,
			Type: "fail",
			Message:  "article does not exist",
		})
	} else {
		//更新
		if (postArticle.Title == "" && postArticle.Content == "" && postArticle.Tag == "") {
			c.JSON(http.StatusOK, &ApiResponse {
				Code: 200,
				Type: "success",
				Message:  "nothing changed",
			})
			return
		}
		if (postArticle.Title == "") {
			postArticle.Title = tmpArticle.Title
		}
		if (postArticle.Content == "") {
			postArticle.Content = tmpArticle.Content
		}
		if (postArticle.Tag == "") {
			postArticle.Tag = tmpArticle.Tag
		}
		MyarticleModel.DB.Update(bson.M{"_id": bson.ObjectIdHex(c.Param("aid"))}, bson.M{"$set": bson.M{
			"title": postArticle.Title,
			"content": postArticle.Content,
			"tag": postArticle.Tag,
			"lastModified": time.Now().Format("2006-01-02 15:04:05"),
		}})
		c.JSON(http.StatusOK, &ApiResponse {
			Code: 200,
			Type: "success",
			Message:  "modify article success",
		})
	}
		
	
}
*/