package hupuAPI

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"log"
	"utils/jex"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/api/players/:round", func(c *gin.Context) {
		round := c.Param("round")
		response, _ := http.Get("http://api.liangle.com/api/passerbyking/game/players/" + round)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		var jo = jex.Load(body)
		var playerArr = jo.GetArray("data")
		for _, player := range playerArr {
			var playerName = player.GetString("name")
			var playerNum = player.GetString("playerNum")
			log.Printf("player %s: %s\n", playerNum, playerName)
		}

		c.JSON(200, jo.Data())
	})

}
