package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//	router.GET("/", func(c *gin.Context) {
	//		c.Redirect(302, "/jungle/")
	//	})

	router.Static("/", "./static/HTML")
	router.Static("/CSS/", "./static/CSS")
	router.Static("/JS/", "./static/JS")
	router.Static("/IMG/", "./static/IMG")

	router.NoRoute(func(c *gin.Context) {
		c.String(404, "Womp Womp, file does not exist")
	})

	//router.LoadHTMLFiles() # for template files later

	router.Run() // listen and serve on 0.0.0.0:8080
}
