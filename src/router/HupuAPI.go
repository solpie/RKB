package hupuAPI

import (

	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"github.com/antonholmquist/jason"
	"log"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/api/players/:round", func(c *gin.Context) {
		round := c.Param("round")
		response, _ := http.Get("http://api.liangle.com/api/passerbyking/game/players/" + round)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		v, _ := jason.NewObjectFromBytes(body)
		//msg, _ := v.GetString("msg")
		//log.Println(msg)
		playerArr, _ := v.GetObjectArray("data")
		for i, player := range playerArr {
			playerName,_ := player.GetString("name")
			log.Printf("player %d: %s", i, playerName)
		}
		c.JSON(200, v)
	})

}
