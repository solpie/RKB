package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"utils"
	"router"
	"fmt"
)

func main() {
	var playerDb = new(godb.GoDB)
	playerDb.Init("./db/player.db")

	fmt.Println(playerDb.Path)
	//fmt.Println(playerDb.DataMap["119"].GetString("name"))

	var router = gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./static/tmpl/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./static/websocket.html")
	})

	router.GET("/admin/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/panel/:name/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "panel.tmpl", gin.H{
			"title": "Main website",
			"wsPort": 6969,
		})
		//name := c.Param("name")
		//message := name
		//c.String(http.StatusOK, message)
	})
	hupuAPI.SetupRouter(router)

	initWS(router)
	//httpTest(router)
	router.Run(":80")
}
func initWS(r *gin.Engine) {
	m := melody.New()
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})
	r.GET("/cmd/:cmdId", func(c *gin.Context) {
		cmdId := c.Param("cmdId")
		m.Broadcast([]byte(cmdId))
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

}
