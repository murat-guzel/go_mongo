package database

import (
	"go_mongo/logger"

	"github.com/globalsign/mgo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
    mgoSession *mgo.Session
    databaseName  = ""
    collection = ""
    address = ""
)



func BuildMongo(_collection string,_databaseName string,_address string){

    collection = _collection
    databaseName =_databaseName
    address = _address
}


//Get Session
func getSession () *mgo.Session {
    if mgoSession == nil {
        var err error
        mgoSession, err = mgo.Dial(address)
        if err != nil {
             panic(err) // no, not really
        }
    }
    return mgoSession.Clone()
}

//Create Connection
func withCollection(s func(*mgo.Collection) error) error {
    session := getSession()
    defer session.Close()
    c := session.DB(databaseName).C(collection)
    return s(c)
}



//MongoSession returns a call to a mongoDB
func MongoSession() *mgo.Session {

	session, err := mgo.Dial("10.132.0.5")
	if err != nil {
		logger.LogError(&err, "test")
	}

	// if err != nil {
	// 	panic(err)
	// }

	return session
}
