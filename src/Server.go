package main

import (
	"net/http"
	"log"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/antonholmquist/jason"
	"github.com/olahol/melody"
)
//https://github.com/gorilla/websocket/blob/master/examples/echo/server.go

func main() {
	router := gin.Default()
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
		})
		//name := c.Param("name")
		//message := name
		//c.String(http.StatusOK, message)
	})
	initWS(router)
	httpTest(router)
	router.Run(":80")
}
func httpTest(r *gin.Engine) {
	r.GET("/api/players/:round", func(c *gin.Context) {
		round := c.Param("round")
		response, _ := http.Get("http://api.liangle.com/api/passerbyking/game/players/" + round)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		v, _ := jason.NewObjectFromBytes(body)
		msg, _ := v.GetString("msg")
		log.Println(msg)
		playerArr, _ := v.GetObjectArray("data")
		for i, player := range playerArr {
			log.Printf("player %d: %s", i, player)
		}
		c.JSON(200, v)
	})

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
