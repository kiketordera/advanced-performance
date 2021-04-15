package repositories

import (
	"fmt"
	"log"

	"github.com/kiketordera/advanced-performance/app/core"
	_domain "github.com/kiketordera/advanced-performance/app/domain"
	"github.com/mailjet/mailjet-apiv3-go"
)

// LoginRepository implements models.ILogin with Custom Login
type MailRepository struct {
}

// InitDatabase creates the instance of the BoltHold Database
func GetCustomMail() _domain.IMail {
	return MailRepository{}
}

func (r MailRepository) SendMailCorfirmation(name, contact, message string) {
	wholeMessage := "Mail from: " + name + ", with contact: " + contact + ", message: " + message
	mailjetClient := mailjet.NewMailjetClient(core.MailjetPublicKey, core.MailjetPrivateKey)
	// MailTo are the recipients of the form
	emailTo := &mailjet.RecipientsV31{
		mailjet.RecipientV31{
			Email: "advdecor@gmail.com",
			// Email: "kiketordera@gmail.com",
			Name: core.KNameProyect,
		},
	}

	mailTeam := []mailjet.InfoMessagesV31{
		mailjet.InfoMessagesV31{
			From:     core.MailFrom,
			To:       emailTo,
			Subject:  core.Subject,
			TextPart: wholeMessage,
			CustomID: core.CustomeID,
		},
	}
	messages := mailjet.MessagesV31{Info: mailTeam}
	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data: %+v\n", res)
	fmt.Println("Email Sent!")
}
