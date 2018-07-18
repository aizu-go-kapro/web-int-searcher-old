package di

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func InjectDB() *mgo.Session {
	var err error

	if session != nil {
		return session
	}
	session, err = mgo.Dial("mongodb://localhost/test")
	if err != nil {
		panic(err)
	}

	return session
}
