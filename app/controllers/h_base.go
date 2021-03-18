package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_domain "github.com/kiketordera/advanced-performance/app/domain"
	_validators "github.com/kiketordera/advanced-performance/app/validators"

	validator "github.com/go-playground/validator/v10"
	i18n "github.com/suisrc/gin-i18n"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	Mail     _domain.IMail
	Validate *validator.Validate
}

// NewBaseHandler returns a new BaseHandler
func NewWebHandler(mail _domain.IMail) *BaseHandler {
	// We add out validators
	validator := validator.New()
	validator.RegisterValidation("my-custom-tag", _validators.MyCustomValidator)
	return &BaseHandler{
		Mail:     mail,
		Validate: validator,
	}
}

// Render renders all templates and export the HTML to the browser the HTML
func (h *BaseHandler) render(c *gin.Context, data gin.H, templateName string) {
	data["GLOBAL"] = h.getGobalTemplateVariables(c)
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
	m["home"] = i18n.FormatMessage(c, &i18n.Message{ID: "home"}, nil)
	m["rent"] = i18n.FormatMessage(c, &i18n.Message{ID: "rent"}, nil)
	m["buy"] = i18n.FormatMessage(c, &i18n.Message{ID: "buy"}, nil)
	m["profile"] = i18n.FormatMessage(c, &i18n.Message{ID: "profile"}, nil)
	m["uploadProperties"] = i18n.FormatMessage(c, &i18n.Message{ID: "uploadProperties"}, nil)
	m["featuredProperties"] = i18n.FormatMessage(c, &i18n.Message{ID: "featuredProperties"}, nil)
	m["contact"] = i18n.FormatMessage(c, &i18n.Message{ID: "contact"}, nil)
	m["register"] = i18n.FormatMessage(c, &i18n.Message{ID: "register"}, nil)
	return m
}

// PrifleDetails shows the page with the details of the profile
func (h *BaseHandler) RenderFeedback(c *gin.Context, title, body, url string, good, goBack bool) {
	h.render(c, gin.H{
		"title": title,
		"body":  body,
		"url":   url,
		"back":  goBack,
		"good":  good,
	}, "feedback.html")
}
