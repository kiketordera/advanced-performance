package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	_domain "github.com/kiketordera/advanced-performance/app/domain"
	_validators "github.com/kiketordera/advanced-performance/app/validators"
	i18n "github.com/suisrc/gin-i18n"

	validator "github.com/go-playground/validator/v10"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	Repository _domain.IRepository
	Login      _domain.ILogin
	Mail       _domain.IMail
	Validate   *validator.Validate
}

// NewBaseHandler returns a new BaseHandler
func NewWebHandler(repo _domain.IRepository, login _domain.ILogin, mail _domain.IMail) *BaseHandler {
	// We add out validators
	validator := validator.New()
	validator.RegisterValidation("my-custom-tag", _validators.MyCustomValidator)
	return &BaseHandler{
		Repository: repo,
		Login:      login,
		Mail:       mail,
		Validate:   validator,
	}
}

// Render renders all templates and export the HTML to the browser the HTML
func (h *BaseHandler) render(c *gin.Context, data gin.H, templateName string) {
	data["GLOBAL"] = h.getGobalTemplateVariables(c)
	fmt.Print("This is global: ", data["GLOBAL"])
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

// This method create the global variables and it passes to all the HTML templates
func (h *BaseHandler) getGobalTemplateVariables(c *gin.Context) map[string]interface{} {
	m := make(map[string]interface{})
	m["about"] = i18n.FormatMessage(c, &i18n.Message{ID: "about_tab_id"}, nil)
	m["services"] = i18n.FormatMessage(c, &i18n.Message{ID: "serv_tab_title_id"}, nil)
	m["become"] = i18n.FormatMessage(c, &i18n.Message{ID: "become"}, nil)
	m["contact"] = i18n.FormatMessage(c, &i18n.Message{ID: "contact_tab_id"}, nil)
	return m
}

// PrifleDetails shows the page with the details of the profile
func (h *BaseHandler) RenderFeedback(c *gin.Context, title, body, url string, good, goBack bool) {
	h.render(c, gin.H{
		"title":     title,
		"body":      body,
		"hideVideo": true,
	}, "feedback.html")
}

// getPhotoFromHTML gets the Photo information of the HTML in a POST request with server validation
// and then saves it in the directory given. If not directory given, it creates it
func GetPhotoFromHTML(c *gin.Context, name string, indexElement string) string {
	directory := "../media/images/uploads/"
	file, err := c.FormFile(name)
	if err != nil {
		fmt.Print("Error: ", err)
		panic("Error retrieving the file from HTML")
	}
	if file.Filename != "" {
		file.Filename = indexElement + path.Ext(file.Filename)
		if err := c.SaveUploadedFile(file, directory+file.Filename); err != nil {
			os.MkdirAll(directory, os.ModePerm)
			if err = c.SaveUploadedFile(file, directory+file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("Upload photo err: %s", err.Error()))
				panic("Upload photo err")
			}
		}
	} else {
		panic("File not retrieved from HTML")
	}
	return file.Filename
}
