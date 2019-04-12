package shared

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

//GetPath return a path of binary
func GetPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
