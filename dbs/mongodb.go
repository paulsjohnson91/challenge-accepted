package dbs

import (
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)


//MgoSession and session
type MgoSession struct {
	Session *mgo.Session
}

func newMgoSession(s *mgo.Session) *MgoSession {
	return &MgoSession{s}
}

//const SERVER = "ds055885.mlab.com:55885"

// DBNAME the name of the DB instance
const DBNAME = "gorest"

// DOCNAME the name of the document
const DOCNAME = "albums"

const AuthUserName = "adminuser"

const AuthPassword = "admpass1"

//StartMongoDB initialize session on mongodb
func StartMongoDB(msg string) *MgoSession {
	log.Info("Attempting to connect to " + os.Getenv("MONGODB_URL"))
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs: []string{os.Getenv("MONGODB_URL")},
		//Addrs:    []string{SERVER},
		Timeout:  60 * time.Second,
		Database: DBNAME,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("[MongoDB] CreateSession: %s\n", err)
	}
	mongoSession.SetMode(mgo.Monotonic, true)

	log.Infof("[MongoDB] connected! %s", msg)
	return newMgoSession(mongoSession)
}
