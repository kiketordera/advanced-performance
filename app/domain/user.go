package domain

import (
	"gopkg.in/mgo.v2/bson"
)

// User Represents our users in the DataBase
type User struct {
	ID       bson.ObjectId
	Name     string
	Email    string
	Password string
	Photos   []string
}

// UserRole is the close list of different roles that the users can have
type UserRole string

// The options for RoleType
const (
	Admin UserRole = "admin"
)
