package main

import "gopkg.in/mgo.v2"

type DbManager struct {
	Database string
}

func NewDbManager() *DbManager{
	return &DbManager{"gameRanking"}
}

func (mgr DbManager) Insert(msg *Result, table string) {
	session, err := mgo.Dial(*DbHost)
	failOnError(err, "Unable to connect to MongoDB")

	c := session.DB(mgr.Database).C(table)
	err = c.Insert(msg)
	failOnError(err, "Unable to insert to database")

	defer session.Close()
}

func (mgr DbManager) Receive(limit int, table string, primarySort string, secondarySort string) []Result {
	session, err := mgo.Dial(*DbHost)
	failOnError(err, "Unable to connect to MongoDB")

	var msg []Result

	c := session.DB(mgr.Database).C(table)
	err = c.Find(nil).Sort(primarySort, secondarySort).Limit(limit).All(&msg)

	failOnError(err, "Unable to select from database")

	defer session.Close()

	return msg
}
