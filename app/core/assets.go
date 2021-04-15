package core

import (
	_domain "github.com/kiketordera/advanced-performance/app/domain"
	"github.com/mailjet/mailjet-apiv3-go"

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

	// MailFrom is the mail address from where we send the mail
	MailFrom = &mailjet.RecipientV31{
		Email: "kiketordera@gmail.com",
		Name:  "The Totem Studio",
	}

	MailTo = &mailjet.RecipientsV31{
		mailjet.RecipientV31{
			Email: "kiketordera@gmail.com",
			Name:  KNameProyect,
		},
	}

	// Subject is the subject of the mail
	Subject = "Mail from TOTEM Studio"

	MailjetPublicKey  = "ef7c3c8d9afc7338cec954515713407d" //ef7
	MailjetPrivateKey = "ab513810d82ba588e64fd2fc871fea4a" //ab
	GoogleMapsAPI     = "AIzaSyC_YNg3llgO1OMqxZCrJ1Nd-g1oN9UpYT4"
	CustomeID         = "AppGettingStartedTest"
)
