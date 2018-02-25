package database

import (
	"go_mongo/logger"

	"github.com/globalsign/mgo"
)

//MongoSession returns a call to a mongoDB
func MongoSession() *mgo.Session {

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		logger.LogError(&err, "test")
	}

	// if err != nil {
	// 	panic(err)
	// }

	return session
}
