package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"router"
	"model"
	"github.com/skip2/go-qrcode"
	"github.com/googollee/go-socket.io"
	"encoding/base64"
	"fmt"
	"log"
	"utils/jex"
	"utils"
)

func InitServer() {
	//gin.SetMode(gin.ReleaseMode)
	//var confJo = jex.LoadFile("./conf.json")
	//fmt.Println("conf.json", confJo.GetNumber("port"))

	model.Db()
	//model.SrvModel().Port = "8082"//string(confJo.GetNumber("port"))

	var ginEngine = gin.Default()
	ginEngine.Static("/static", "./static")
	ginEngine.LoadHTMLGlob("./static/tmpl/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	// This handler will match /user/john but will not match neither /user/ or /user
	ginEngine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin")
		//http.ServeFile(c.Writer, c.Request, "./static/websocket.html")
	})

	ginEngine.GET("/qrcode", func(c *gin.Context) {
		//var png []byte
		var png, _ = qrcode.Encode("https://example.org", qrcode.Medium, 256)
		//base64Text := make([]byte, 0)
		//base64.StdEncoding.Encode(base64Text, png)
		str := base64.StdEncoding.EncodeToString(png)
		var imgSrc = "data:image/png;base64," + str
		fmt.Println(str)
		c.String(200, imgSrc)
		//c.String(200, `<img str="` + str + `">`)
		//base64.NewEncoder(base64.StdEncoding,png)
	})

	ginEngine.GET("/panel/:name/", func(c *gin.Context) {
		name := c.Param("name")
		c.HTML(http.StatusOK, "panel.tmpl", gin.H{
			"panelName": name,
		})
		//name := c.Param("name")
		//message := name
		//c.String(http.StatusOK, message)
	})
	router.SetupHupuAPI(ginEngine)
	router.SetupDb(ginEngine)

	initWS(ginEngine)
	ginEngine.Run(":80")
	//ginEngine.Run(":" + model.SrvModel().Port)
}

func initWS(r *gin.Engine) {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		so.Join("panel")
		//so.On("chat message", func(msg string) {
		//	log.Println("emit:", so.Emit("chat message", msg))
		//	so.BroadcastTo("chat", "chat message", msg)
		//})
		so.On("opUrl", func(msg utils.JParam) {
			var jo = jex.Load(msg.JsonStr)
			var opUrl = jo.GetString("opUrl")
			//model.SrvModel().OpUrlMap[opUrl] = opUrl

			log.Println("opUrl", opUrl)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	r.GET("/socket.io/", func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
		//m.HandleRequest(c.Writer, c.Request)
	})

	r.GET("/cmd/:cmdId", func(c *gin.Context) {
		cmdId := c.Param("cmdId")
		log.Println(cmdId, c.Request.Body)
	})

	r.POST("/cmd/:cmdId", func(c *gin.Context) {
		var url utils.JParam
		c.BindJSON(&url)
		var jo = jex.Load(url.JsonStr)
		var b = jo.GetBool("bool")
		var n = jo.GetNumber("num")
		var s = jo.GetString("string")
		log.Println("PostForm", b, s, n)

		c.String(200, "ok")
	})
}
