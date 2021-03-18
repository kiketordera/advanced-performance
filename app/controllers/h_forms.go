package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
This method get the form from the HTML, and makes a safe validation of the camps
*/
func (h *BaseHandler) getForm(c *gin.Context) {
	// Safe validation with anonymous struct
	formData := &struct {
		Name    string `form:"fname" binding:"required" san:"max=50,trim,title,xss"`
		Contact string `form:"fcontact" binding:"required" san:"max=35,trim,xss"`
		Message string `form:"fmessage" binding:"required" san:"max=1250,trim,xss"`
	}{}
	validateSanitaze(c, formData)
	fmt.Print("This is the struct: ", formData)
	c.Redirect(http.StatusFound, "#good-feedback")
}

// validateSanitaze validates and sanitizes the inputs to avoid code injection
func (h *BaseHandler) validateSanitaze(c *gin.Context, st interface{}) {
	// Validation (with Gin)
	if err := c.Bind(st); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Sanitization (with Sanitizer)
	s, err := sanitizer.New()
	if err != nil {
		fmt.Print("Error: ", err)
		panic("It has beeen an error sanitizing the contact form: ")
	}
	err = s.Sanitize(st)
	if err != nil {
		fmt.Print("Error: ", err)
		panic("It has beeen an error sanitizing the contact form: ")
	}
}
