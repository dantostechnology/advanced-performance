package main

import (
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
	// We create the instance for Gin
	r := gin.Default()

	// Internationalization for showing the right language to match the browser's  default settings
	// The name of the files must match the
	bundle := i18n.NewBundle(
		language.English,
		"text/en.toml",
		"text/es.toml",
	)

	// Tell Gin to use our middleware. This means that in every single request (GET, POST...), the call to i18n will be executed
	r.Use(i18n.Serve(bundle))

	// Path to the static files (images, svg, css..)
	r.StaticFS("/static", http.Dir("/media"))

	// Path to the HTML templates
	r.LoadHTMLGlob("/*.html")

	// Redirects when users introduces a wrong URL
	r.NoRoute(redirect)

	// This get executed when the users gets into our website in the home domain ("/")
	r.GET("/", renderHome)
	// Listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}

/* Renders the landing page and it passes the parameters that will be rendered in the HTML.
In this case the text of the website, and we are using the i18n to detect the default browser language of the user and show accordingly.
*/
func renderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"hi": i18n.FormatMessage(c, &i18n.Message{ID: "hi"}, nil),
	})
}

// Redirects to the home route when the users type an URL inside our domain that does not exists
func redirect(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
