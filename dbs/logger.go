package dbs

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
)

//Logger hook
func Logger() *logrus.Logger {
	logger := logrus.New()
	hooker, err := mgorus.NewHookerWithAuth("ds055885.mlab.com:55885", "gorest", "logs", "adminuser", "admpass1")
	if err == nil {
		logger.Hooks.Add(hooker)
	} else {
		fmt.Print(err)
	}

	return logger
}
