package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	i18n "github.com/suisrc/gin-i18n"
)

/* Renders the landing page and it passes the parameters that will be rendered in the HTML.
In this case the text of the website, and we are using the i18n to detect the default browser language of the user and show accordingly.
*/
func (h *BaseHandler) RenderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "landing.html", gin.H{
		"hi": i18n.FormatMessage(c, &i18n.Message{ID: "hi"}, nil),
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
