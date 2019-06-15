package DataBase

import(
	"gopkg.in/mgo.v2"
	_"gopkg.in/mgo.v2/bson"
)

func Conn() (db *mgo.Database) {
	session,err := mgo.Dial("localhost:27017") //传url，连接数据库
	if err != nil {
		return
	}
	defer session.Close()
	//session.SetMode(mgo.Monotonic, true)
	db = session.DB("test")
	
	return
}