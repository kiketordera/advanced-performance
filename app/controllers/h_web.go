package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_core "github.com/kiketordera/advanced-performance/app/core"
	_domain "github.com/kiketordera/advanced-performance/app/domain"
	i18n "github.com/suisrc/gin-i18n"
	"golang.org/x/crypto/bcrypt"
)

/* Renders the landing page and it passes the parameters that will be rendered in the HTML.
In this case the text of the website, and we are using the i18n to detect the default browser language of the user and show accordingly.
*/
func (h *BaseHandler) RenderHome(c *gin.Context) {
	u := h.Repository.GetAdminUser()
	c.HTML(http.StatusOK, "landing.html", gin.H{
		"hi":   i18n.FormatMessage(c, &i18n.Message{ID: "hi"}, nil),
		"user": u,
	})
}

/* Renders the dealers page and it passes the parameters that will be rendered in the HTML.
In this case the text of the website, and we are using the i18n to detect the default browser language of the user and show accordingly.
*/
func (h *BaseHandler) RenderDealers(c *gin.Context) {
	c.HTML(http.StatusOK, "dealers.html", gin.H{
		"hi": i18n.FormatMessage(c, &i18n.Message{ID: "hi"}, nil),
	})
}

/* Renders the dealers page and it passes the parameters that will be rendered in the HTML.
In this case the text of the website, and we are using the i18n to detect the default browser language of the user and show accordingly.
*/
func (h *BaseHandler) RenderLogin(c *gin.Context) {
	fmt.Print("Is user logged in: ", h.Login.IsLoggedIn(c))
	c.HTML(http.StatusOK, "login.html", gin.H{
		"hideVideo": true,
	})
}

func (h *BaseHandler) RenderDashboard(c *gin.Context) {
	u := h.Repository.GetAdminUser()
	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"user":      u,
		"hideVideo": true,
	})
}

func (h *BaseHandler) RenderEditPhoto(c *gin.Context) {
	i := c.Param("index")
	index, _ := strconv.Atoi(i)
	c.HTML(http.StatusOK, "upload-image.html", gin.H{
		"hideVideo": true,
		"index":     index + 1,
	})
}

func (h *BaseHandler) EditPhoto(c *gin.Context) {
	i := c.Param("index")
	index, _ := strconv.Atoi(i)
	name := GetPhotoFromHTML(c, "photo", i)
	u := h.Repository.GetAdminUser()
	u.Photos[index] = name
	h.Repository.SaveObject(u, u.ID)
	c.Redirect(http.StatusFound, "/dashboard")
}

// GetLoginForm gets the params from the login form in the HTML, and makes a safe validation of the fields
func (h *BaseHandler) GetLoginForm(c *gin.Context) {
	// Safe sanitization and validation with anonymous struct
	formData := &struct {
		Email    string `form:"femail" binding:"required" san:"max=25,trim,lower,xss"`
		Password string `form:"fpassword" binding:"required" san:"max=50,trim,xss"`
	}{}
	_core.ValidateSanitaze(c, formData, h.Validate)
	u, b := h.Repository.GetUserByMail(formData.Email)
	if !b {
		h.RenderFeedback(c, "Wrong mail", "The mail introduced is not valid", "/", false, true)
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(formData.Password))
	if err == nil {
		fmt.Print("Pues estro deberia poner la coookkie")
		h.Login.SetSession(c, u.Email, _domain.Admin)
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}
	h.RenderFeedback(c, "Wrong password", "The password introduced is not valid", "/", false, true)
}
