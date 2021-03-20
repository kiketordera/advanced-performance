package core

import (
	_domain "github.com/kiketordera/advanced-performance/app/domain"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

const (

	// KNameProyect is the name of the proyect
	KNameProyect = "advanced-performance"

	// KPort is the port where the server will be running
	KPort = ":8050"

	// SessionTime is the seconds the session will be active
	SessionTime = 108000
)

var (
	// TokenSigningKey is the key to sign the cookies
	TokenSigningKey   = []byte("SuperFancyToken:D")
	KPasswordAdmin, _ = bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	KAdminUser        = _domain.User{
		ID:       bson.NewObjectId(),
		Name:     "Admin",
		Email:    "admin@admin",
		Password: string(KPasswordAdmin),
		Photos:   []string{"0.png", "1.png", "2.png", "3.png", "4.png", "5.png"},
	}
)
