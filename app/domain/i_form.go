package domain

// ILogin are all the Usecases of the login
type IMail interface {
	SendMailCorfirmation(name, contact, message string)
}
