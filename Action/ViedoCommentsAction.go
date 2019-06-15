package Action

import(
	db "VideoAPI/Resources/DataBase"
	"VideoAPI/Models"
	"gopkg.in/mgo.v2/bson"
	"VideoAPI/Resources"
	"time"
)

func GetVideoCommentsAction(vedioid string) (ret map[string]interface{}, err error) {

	conn := db.Conn()
	collection := conn.C("comments")
	
	comments := Models.Comments{}
	err = collection.Find(bson.M{"VIDEOID":vedioid}).All(&comments.CommentList)
	if err != nil {
		return
	}

	ret,err = Resources.StructToMapJson(comments)
	return
}

func AddVideoCommentAction(c Models.Comment) (err error) {

	conn := db.Conn()
	collection := conn.C("comments")

	c.CommentTime = time.Now()

	err = collection.Insert(&c)
	if err != nil {
		return
	}
	
	return
}