package services

import (
	log "github.com/sirupsen/logrus"

	"gopkg.in/mgo.v2/bson"

	"fmt"

	db "github.com/paulsjohnson91/challenge-accepted/dbs"
	model "github.com/paulsjohnson91/challenge-accepted/models"
	lib "github.com/paulsjohnson91/challenge-accepted/shared"
)

var session = db.StartMongoDB("Middleware / User Service").Session

//UserIsValidOnProject seek the user on the project profile
func UserIsValidOnProject(slug string, userID string) (model.Project, error) {

	ss := session.Copy()
	defer ss.Close()

	oid := bson.ObjectIdHex(userID)

	//find user
	u := model.Project{}
	if err := ss.DB("gorest").C("projects").Find(bson.M{"slug": slug, "users": oid}).One(&u); err != nil {
		return u, fmt.Errorf("User not valid on project")
	}

	return u, nil
}

// UserGetPermissions return a permisson of user by project and endpoint
func UserGetPermissions(userID string, projectID string, endpoint string) (model.Permission, error) {

	ss := session.Copy()
	defer ss.Close()

	//change objectId
	endp := lib.GetRootEndpoint(endpoint)
	oid := bson.ObjectIdHex(userID)
	oidp := bson.ObjectIdHex(projectID)

	//find user
	u := model.Permission{}
	if err := ss.DB("gorest").C("permissions").Find(bson.M{"owner": oid, "project": oidp, "endpoint": endp}).One(&u); err != nil {
		return u, fmt.Errorf("Permission not found")
	}

	return u, nil
}

func UserAdminStatus(id string) bool {

	ss := session.Copy()
	defer ss.Close()

	//find user
	u := model.User{}
	if !bson.IsObjectIdHex(id) {
		log.Println("ID is not BSON ID")
		return false
	}

	oid := bson.ObjectIdHex(id)
	if err := ss.DB("gorest").C("users").FindId(oid).One(&u); err != nil {
		return false
	}
	return u.Admin

}
