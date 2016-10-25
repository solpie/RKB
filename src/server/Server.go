package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"router"
	"model"
)

func InitServer() {
	//gin.SetMode(gin.ReleaseMode)
	model.Db()
	var ginEngine = gin.Default()
	ginEngine.Static("/static", "./static")
	ginEngine.LoadHTMLGlob("./static/tmpl/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	// This handler will match /user/john but will not match neither /user/ or /user
	ginEngine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin")
		//http.ServeFile(c.Writer, c.Request, "./static/websocket.html")
	})

	ginEngine.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.tmpl", gin.H{
			"version": "0.4",
		})
	})

	ginEngine.GET("/panel/:name/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "panel.tmpl", gin.H{
			"title": "Main website",
			"wsPort": 6969,
		})
		//name := c.Param("name")
		//message := name
		//c.String(http.StatusOK, message)
	})
	router.SetupHupuAPI(ginEngine)
	router.SetupDb(ginEngine)

	initWS(ginEngine)
	//httpTest(router)
	ginEngine.Run(":80")
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
