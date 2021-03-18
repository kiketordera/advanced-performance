package repositories

import (
	_domain "github.com/kiketordera/advanced-performance/app/domain"
)

// LoginRepository implements models.ILogin with Custom Login
type MailRepository struct {
}

// InitDatabase creates the instance of the BoltHold Database
func GetCustomMail() _domain.IMail {
	return MailRepository{}
}

func (r MailRepository) SendMailCorfirmation() {
	//TODO: implement mailing service
}
