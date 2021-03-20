package domain

import (
	"gopkg.in/mgo.v2/bson"
)

// IRepository are all the Usecases of the database
type IRepository interface {
	SaveObject(o interface{}, id bson.ObjectId) error
	GetAdminUser() User
	GetUserByMail(email string) (User, bool)
}
