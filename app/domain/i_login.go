package domain

import (
	"github.com/gin-gonic/gin"
)

// ILogin are all the Usecases of the login
type ILogin interface {
	CheckToken() gin.HandlerFunc
	SetSession(c *gin.Context, username string, role UserRole)
	EnsureLoggedIn(c *gin.Context)
	IsLoggedIn(c *gin.Context) bool
	EnsureNotLoggedIn(c *gin.Context)
	Logout(c *gin.Context)
	EnsureIsAdmin(c *gin.Context)
}
