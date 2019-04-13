package dbs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"github.com/paulsjohnson91/challenge-accepted/logger"
)

	var log = logger.Logger()

//Dispatch choose a db session
// add new dispath of other database just put here
// the session.
type Dispatch struct {
	MongoDB *mgo.Session
	Logger  *logrus.Logger
}

//StartDispatch load up connections
func StartDispatch() *Dispatch {
	fmt.Println("Starting db connection")
	//add session of mongodb
	mongosession := StartMongoDB("Dispatch Service").Session
	// add logger for dispatch

	return &Dispatch{MongoDB: mongosession, Logger: logger.Logger()}

}
