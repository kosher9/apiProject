package dao

import (
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/mgo.v2"
	"log"
	. "user/apiProject/models"
)

type UserDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const COLLECTION = "users"

// Find all users
func (u *UserDAO) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// Establish a connection to database
func (u *UserDAO) Connect() {
	session, err := mgo.Dial(u.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(u.Database)
}

/*func CreatePerson(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("content-type", "application/json")
	var user models.User
}
*/
