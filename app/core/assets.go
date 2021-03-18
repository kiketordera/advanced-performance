package core

const (

	// KNameProyect is the name of the proyect
	KNameProyect = "advanced-performance"

	// KPort is the port where the server will be running
	KPort = ":8050"
)

var (
	// TokenSigningKey is the key to sign the cookies
	TokenSigningKey = []byte("SuperFancyToken:D")
)
