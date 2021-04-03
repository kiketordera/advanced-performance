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
	h.render(c, gin.H{
		"about_main_title_id":  i18n.FormatMessage(c, &i18n.Message{ID: "about_main_title_id"}, nil),
		"about_main_text1_id":  i18n.FormatMessage(c, &i18n.Message{ID: "about_main_text1_id"}, nil),
		"about_main_text2_id":  i18n.FormatMessage(c, &i18n.Message{ID: "about_main_text2_id"}, nil),
		"your_id":              i18n.FormatMessage(c, &i18n.Message{ID: "your_id"}, nil),
		"your_text_id":         i18n.FormatMessage(c, &i18n.Message{ID: "your_text_id"}, nil),
		"works_id":             i18n.FormatMessage(c, &i18n.Message{ID: "works_id"}, nil),
		"works_text_id":        i18n.FormatMessage(c, &i18n.Message{ID: "works_text_id"}, nil),
		"works_text2_id":       i18n.FormatMessage(c, &i18n.Message{ID: "works_text2_id"}, nil),
		"install_id":           i18n.FormatMessage(c, &i18n.Message{ID: "install_id"}, nil),
		"install_text_id":      i18n.FormatMessage(c, &i18n.Message{ID: "install_text_id"}, nil),
		"install_text2_id":     i18n.FormatMessage(c, &i18n.Message{ID: "install_text2_id"}, nil),
		"serv_tab_title_id":    i18n.FormatMessage(c, &i18n.Message{ID: "serv_tab_title_id"}, nil),
		"file_title_id":        i18n.FormatMessage(c, &i18n.Message{ID: "file_title_id"}, nil),
		"file_text_id":         i18n.FormatMessage(c, &i18n.Message{ID: "file_text_id"}, nil),
		"soft_optim_title_id":  i18n.FormatMessage(c, &i18n.Message{ID: "soft_optim_title_id"}, nil),
		"soft_optim_text1_id":  i18n.FormatMessage(c, &i18n.Message{ID: "soft_optim_text1_id"}, nil),
		"soft_optim_text2_id":  i18n.FormatMessage(c, &i18n.Message{ID: "soft_optim_text2_id"}, nil),
		"dyno_title_id":        i18n.FormatMessage(c, &i18n.Message{ID: "dyno_title_id"}, nil),
		"dyno_text1_id":        i18n.FormatMessage(c, &i18n.Message{ID: "dyno_text1_id"}, nil),
		"dyno_text2_id":        i18n.FormatMessage(c, &i18n.Message{ID: "dyno_text2_id"}, nil),
		"device_inst_title_id": i18n.FormatMessage(c, &i18n.Message{ID: "device_inst_title_id"}, nil),
		"device_inst_text1_id": i18n.FormatMessage(c, &i18n.Message{ID: "device_inst_text1_id"}, nil),
		"device_inst_text2_id": i18n.FormatMessage(c, &i18n.Message{ID: "device_inst_text2_id"}, nil),
		"device_inst_text3_id": i18n.FormatMessage(c, &i18n.Message{ID: "device_inst_text3_id"}, nil),
		"user":                 u,
	}, "landing.html")
}

/* Renders the dealers page and it passes the parameters that will be rendered in the HTML.
In this case the text of the website, and we are using the i18n to detect the default browser language of the user and show accordingly.
*/
func (h *BaseHandler) RenderDealers(c *gin.Context) {
	u := h.Repository.GetAdminUser()
	h.render(c, gin.H{
		"why_title_id":       i18n.FormatMessage(c, &i18n.Message{ID: "why_title_id"}, nil),
		"why_text1_id":       i18n.FormatMessage(c, &i18n.Message{ID: "why_text1_id"}, nil),
		"why_text2_id":       i18n.FormatMessage(c, &i18n.Message{ID: "why_text2_id"}, nil),
		"why_text3_id":       i18n.FormatMessage(c, &i18n.Message{ID: "why_text3_id"}, nil),
		"become_tab_id":      i18n.FormatMessage(c, &i18n.Message{ID: "become_tab_id"}, nil),
		"buy_title_id":       i18n.FormatMessage(c, &i18n.Message{ID: "buy_title_id"}, nil),
		"buy_step1_title_id": i18n.FormatMessage(c, &i18n.Message{ID: "buy_step1_title_id"}, nil),
		"buy_step1_text_id":  i18n.FormatMessage(c, &i18n.Message{ID: "buy_step1_text_id"}, nil),
		"buy_step2_title_id": i18n.FormatMessage(c, &i18n.Message{ID: "buy_step2_title_id"}, nil),
		"buy_step2_text_id":  i18n.FormatMessage(c, &i18n.Message{ID: "buy_step2_text_id"}, nil),
		"buy_step3_title_id": i18n.FormatMessage(c, &i18n.Message{ID: "buy_step3_title_id"}, nil),
		"buy_step3_text_id":  i18n.FormatMessage(c, &i18n.Message{ID: "buy_step3_text_id"}, nil),
		"buy_step4_title_id": i18n.FormatMessage(c, &i18n.Message{ID: "buy_step4_title_id"}, nil),
		"buy_step4_text_id":  i18n.FormatMessage(c, &i18n.Message{ID: "buy_step4_text_id"}, nil),
		"price_title_id":     i18n.FormatMessage(c, &i18n.Message{ID: "price_title_id"}, nil),
		"contains":           i18n.FormatMessage(c, &i18n.Message{ID: "contains"}, nil),
		"pack1_title_id":     i18n.FormatMessage(c, &i18n.Message{ID: "pack1_title_id"}, nil),
		"pack1_te1_id":       i18n.FormatMessage(c, &i18n.Message{ID: "pack1_te1_id"}, nil),
		"pack1_te2_id":       i18n.FormatMessage(c, &i18n.Message{ID: "pack1_te2_id"}, nil),
		"pack1_te3_id":       i18n.FormatMessage(c, &i18n.Message{ID: "pack1_te3_id"}, nil),
		"pack1_te4_id":       i18n.FormatMessage(c, &i18n.Message{ID: "pack1_te4_id"}, nil),
		"pack2_title_id":     i18n.FormatMessage(c, &i18n.Message{ID: "pack2_title_id"}, nil),
		"pack2_t1_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack2_t1_id"}, nil),
		"pack2_t2_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack2_t2_id"}, nil),
		"pack2_t3_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack2_t3_id"}, nil),
		"pack2_t4_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack2_t4_id"}, nil),
		"pack3_title_id":     i18n.FormatMessage(c, &i18n.Message{ID: "pack3_title_id"}, nil),
		"pack3_t1_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack3_t1_id"}, nil),
		"pack3_t2_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack3_t2_id"}, nil),
		"pack4_title_id":     i18n.FormatMessage(c, &i18n.Message{ID: "pack4_title_id"}, nil),
		"pack4_t1_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack4_t1_id"}, nil),
		"pack4_t2_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack4_t2_id"}, nil),
		"pack4_t3_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack4_t3_id"}, nil),
		"pack4_t4_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack4_t4_id"}, nil),
		"pack5_title_id":     i18n.FormatMessage(c, &i18n.Message{ID: "pack5_title_id"}, nil),
		"pack5_t1_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack5_t1_id"}, nil),
		"pack5_t2_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack5_t2_id"}, nil),
		"pack5_t3_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack5_t3_id"}, nil),
		"pack5_t4_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack5_t4_id"}, nil),
		"pack5_t5_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack5_t5_id"}, nil),
		"pack5_t6_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack5_t6_id"}, nil),
		"pack6_title_id":     i18n.FormatMessage(c, &i18n.Message{ID: "pack6_title_id"}, nil),
		"pack6_t1_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack6_t1_id"}, nil),
		"pack6_t2_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack6_t2_id"}, nil),
		"pack6_t3_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack6_t3_id"}, nil),
		"pack6_t4_id":        i18n.FormatMessage(c, &i18n.Message{ID: "pack6_t4_id"}, nil),
		"contact_tab_id":     i18n.FormatMessage(c, &i18n.Message{ID: "contact_tab_id"}, nil),
		"become_text_id":     i18n.FormatMessage(c, &i18n.Message{ID: "become_text_id"}, nil),
		"form_title_id":      i18n.FormatMessage(c, &i18n.Message{ID: "form_title_id"}, nil),
		"form_1_id":          i18n.FormatMessage(c, &i18n.Message{ID: "form_1_id"}, nil),
		"form_2_id":          i18n.FormatMessage(c, &i18n.Message{ID: "form_2_id"}, nil),
		"form_3_id":          i18n.FormatMessage(c, &i18n.Message{ID: "form_3_id"}, nil),
		"contact_id":         i18n.FormatMessage(c, &i18n.Message{ID: "contact_id"}, nil),
		"insta_clients":      i18n.FormatMessage(c, &i18n.Message{ID: "insta_clients"}, nil),
		"insta_id":           i18n.FormatMessage(c, &i18n.Message{ID: "insta_id"}, nil),
		"become":             i18n.FormatMessage(c, &i18n.Message{ID: "become"}, nil),
		"user":               u,
	}, "dealers.html")
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
		h.Login.SetSession(c, u.Email, _domain.Admin)
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}
	h.RenderFeedback(c, "Wrong password", "The password introduced is not valid", "/", false, true)
}
