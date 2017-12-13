package common

import (
	"fmt"
	"work-codes/gohome/app/config"

	"gopkg.in/mgo.v2"
)

func init() {
	fmt.Println("MGO_URL: ", config.DBConfig.Url)
}

// const (
// 	MONGO_URL = "mongodb://127.0.0.1:27017"
// )

func mgoSession() *mgo.Session {
	session, err := mgo.Dial(config.DBConfig.Url)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func DB(dbname string) *mgo.Database {
	db := mgoSession().DB(dbname)
	return db
}

func MgoClose() {
	mgoSession().Close()
}
