package util

import (
	"fmt"

	"github.com/globalsign/mgo"
)

const server = "mongodb://database:27017"
const dbName = "avtaldb"

func OpenSession() *mgo.Session {
	session, err := mgo.Dial(server)
	if err != nil {
		fmt.Println("Failed to establish connection to server.", err)
	}
	return session
}

func FetchCollection(collectionName string) (*mgo.Collection, *mgo.Session) {
	session := OpenSession()
	c := session.DB(dbName).C(collectionName)
	return c, session
}

func LogCorruptData(err error) {
	fmt.Println("Database data is corrupt.", err)
}
