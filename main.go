package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	i18n "github.com/suisrc/gin-i18n"
	"golang.org/x/text/language"
)

var (
	// BasePath is the path to the project
	BasePath = os.Getenv("GOPATH") + "/src/github.com/kiketordera/advanced-performance"
)

func main() {
	r := gin.Default()

	// Internationalization
	bundle := i18n.NewBundle(
		language.Spanish,
		"text/en.toml",
		"text/es.toml",
	)
	r.Use(i18n.Serve(bundle))

	// Path to the static files (images, svg, css..)
	r.StaticFS("/static", http.Dir(BasePath+"/media"))
	// Path to the HTML templates
	r.LoadHTMLGlob(BasePath + "/*.html")

	// Redirects when users introduces a wrong URL
	r.NoRoute(redirect)

	r.GET("/", renderHome)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// Renders the landing page
func renderHome(c *gin.Context) {
	l := parseTags(c.Request.FormValue("lang"), c.Request.Header.Get("Accept-Language"))
	fmt.Print("L es: ", l)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"hi": i18n.FormatMessage(c, &i18n.Message{ID: "hi"}, nil),
	})
}

// Redirects to the home route
func redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}

// Return the Language tag from the browser info
func parseTags(langs ...string) []language.Tag {
	tags := []language.Tag{}
	for _, lang := range langs {
		t, _, err := language.ParseAcceptLanguage(lang)
		if err != nil {
			continue
		}
		tags = append(tags, t...)
	}
	return tags
}
