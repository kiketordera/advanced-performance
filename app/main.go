package main

import (
	"fmt"
	"net/http"

	_controllers "github.com/kiketordera/advanced-performance/app/controllers"
	_core "github.com/kiketordera/advanced-performance/app/core"
	_repository "github.com/kiketordera/advanced-performance/app/repository"

	"github.com/gin-gonic/gin"
	sanitizer "github.com/go-sanitize/sanitize"
	i18n "github.com/suisrc/gin-i18n"
	"golang.org/x/text/language"
)

func main() {

	// Get the repository of the mail
	mail := _repository.GetCustomMail()

	webHandler := _controllers.NewWebHandler(mail)

	// Init the router, listen and serve on 0.0.0.0:8040 (for windows "localhost:8040")
	CreateInterfaceRouter(webHandler).Run(_core.KPort)
}

func CreateInterfaceRouter(h *_controllers.BaseHandler) *gin.Engine {

	// We create the instance for Gin
	r := gin.Default()

	/* Internationalization for showing the right language to match the browser's  default settings

	The name of the files must match:
	af, ar, az, be, bg, bn, bs, ca, cs, cy, da, de, de-AT, de-CH, de-DE, el, el-CY, en, en-AU, en-CA, en-GB,
	en-IE, en-IN, en-NZ, en-US, en-ZA, en-CY, en-TT, eo, es, es-419, es-AR, es-CL, es-CO, es-CR, es-EC, es-ES,
	es-MX, es-NI, es-PA, es-PE, es-US, es-VE, et, eu, fa, fi, fr, fr-CA, fr-CH, fr-FR, gl, he, hi, hi-IN, hr,
	hu, id, is, it, it-CH, ja, ka, km, kn, ko, lb, lo, lt, lv, mk, ml, mn, mr-IN, ms, nb, ne, nl, nn, oc, or,
	pa, pl, pt, pt-BR, rm, ro, ru, sk, sl, sq, sr, st, sw, ta, te, th, tl, tr, tt, ug, ur, uz, vi, wo, zh-CN,
	zh-HK, zh-TW, zh-YUE
	*/
	bundle := i18n.NewBundle(
		language.English,
		"../media/text/en.toml",
		"../media/text/es.toml",
	)

	// Tell Gin to use our middleware. This means that in every single request (GET, POST...), the call to i18n will be executed
	r.Use(i18n.Serve(bundle))

	// Path to the static files. /static is rendered in the HTML and /media is the link to the path to the  images, svg, css.. the static files
	r.StaticFS("/static", http.Dir("../media"))

	// Path to the HTML templates. * is a wildcard
	r.LoadHTMLGlob("../media/html/*/*.html")

	// Redirects when users introduces a wrong URL
	r.NoRoute(redirect)

	// This get executed when the users gets into our website in the home domain ("/")
	r.GET("/", renderHome)
	r.POST("/", getForm)

	r.GET("/dealers", renderDealers)
	r.POST("/dealers", getForm)
	return r
}
