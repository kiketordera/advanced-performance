package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_core "github.com/kiketordera/advanced-performance/app/core"
)

/*
This method get the form from the HTML, and makes a safe validation of the camps
*/
func (h *BaseHandler) GetForm(c *gin.Context) {
	// Safe validation with anonymous struct
	formData := &struct {
		Name    string `form:"fname" binding:"required" san:"max=50,trim,title,xss"`
		Contact string `form:"fcontact" binding:"required" san:"max=35,trim,xss"`
		Message string `form:"fmessage" binding:"required" san:"max=1250,trim,xss"`
	}{}
	_core.ValidateSanitaze(c, formData, h.Validate)
	h.Mail.SendMailCorfirmation(formData.Name, formData.Contact, formData.Message)
	c.Redirect(http.StatusFound, "/dealers/#good-feedback")
}
