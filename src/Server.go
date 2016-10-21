package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	router := gin.Default()
	router.Static("/static", "./static")

	router.LoadHTMLGlob("./static/tmpl/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//router.GET("/panel", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "panel.tmpl", gin.H{
	//		"title": "Main website",
	//	})
	//})
	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/admin/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/panel/:name/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "panel.tmpl", gin.H{
			"title": "Main website",
		})
		//name := c.Param("name")
		//message := name
		//c.String(http.StatusOK, message)
	})

	router.Run(":80")
}
