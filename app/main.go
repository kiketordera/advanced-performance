package main

import (
	"net/http"

	_controllers "github.com/kiketordera/advanced-performance/app/controllers"
	_core "github.com/kiketordera/advanced-performance/app/core"
	_repository "github.com/kiketordera/advanced-performance/app/repositories"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	i18n "github.com/suisrc/gin-i18n"
	"golang.org/x/text/language"
)

func main() {
	// Get the repository we want acording to the DataBase to use
	repo := _repository.InitBoltDatabase()

	// Get the Login handler according with the Login that we want
	login := _repository.GetCustomLogin()

	// Get the repository of the mail
	mail := _repository.GetCustomMail()

	webHandler := _controllers.NewWebHandler(repo, login, mail)

	// Init the router, listen and serve on 0.0.0.0:8040 (for windows "localhost:8040")
	CreateInterfaceRouter(webHandler).Run(_core.KPort)
}

func CreateInterfaceRouter(h *_controllers.BaseHandler) *gin.Engine {

	// We create the instance for Gin
	r := gin.Default()

	// Internationalization for showing the right language to match the browser's  default settings
	bundle := i18n.NewBundle(
		language.English,
		"../media/text/en.toml",
		"../media/text/es.toml",
	)

	// Tell Gin to use our middleware. This means that in every single request (GET, POST...), the call to i18n will be executed
	r.Use(_core.SecureMiddleware(), h.Login.CheckToken(), i18n.Serve(bundle))

	// Path to the static files. /static is rendered in the HTML and /media is the link to the path to the  images, svg, css.. the static files
	r.StaticFS("/static", http.Dir("../media"))

	// Path to the HTML templates. * is a wildcard
	r.LoadHTMLGlob("../media/html/*/*.html")

	// Redirects when users introduces a wrong URL
	r.NoRoute(redirect)

	// This is the route where we will see the documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// This get executed when the users gets into our website in the home domain ("/")
	r.GET("/", h.RenderHome)
	r.POST("/", h.GetForm)

	r.GET("/dealers", h.RenderDealers)
	r.POST("/dealers", h.GetForm)

	r.GET("/login", h.Login.EnsureNotLoggedIn, h.RenderLogin)
	r.POST("/login", h.Login.EnsureNotLoggedIn, h.GetLoginForm)

	r.GET("/dashboard", h.Login.EnsureLoggedIn, h.RenderDashboard)

	r.GET("/edit-photo/:index", h.Login.EnsureLoggedIn, h.RenderEditPhoto)
	r.POST("/edit-photo/:index", h.Login.EnsureLoggedIn, h.EditPhoto)

	r.GET("/logout", h.Login.Logout)

	return r
}

// Redirects to the home route when the users type an URL inside our domain that does not exists
func redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
