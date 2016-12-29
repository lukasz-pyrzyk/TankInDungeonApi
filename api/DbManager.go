package main

import "gopkg.in/mgo.v2"

type DbManager struct {
	Database string
	Table string
}

func NewDbManager() *DbManager{
	return &DbManager{"gameRanking", "results"}
}

func (mgr DbManager) Insert(msg *Result) {
	session, err := mgo.Dial(*DbHost)
	failOnError(err, "Unable to connect to MongoDB")

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mgr.Database).C(mgr.Table)
	err = c.Insert(msg)
	failOnError(err, "Unable to insert to database")

	defer session.Close()
}

func (mgr DbManager) Receive(limit int) []Result {
	session, err := mgo.Dial(*DbHost)
	failOnError(err, "Unable to connect to MongoDB")

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	var msg []Result

	c := session.DB(mgr.Database).C(mgr.Table)
	err = c.Find(nil).Sort("-score", "time").Limit(limit).All(&msg)

	failOnError(err, "Unable to select from database")

	defer session.Close()

	return msg
}
