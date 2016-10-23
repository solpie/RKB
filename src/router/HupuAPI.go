package hupuAPI

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"log"
	"github.com/Jeffail/gabs"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/api/players/:round", func(c *gin.Context) {
		round := c.Param("round")
		response, _ := http.Get("http://api.liangle.com/api/passerbyking/game/players/" + round)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		v, _ := gabs.ParseJSON(body)
		playerArr, _ := v.S("data").Children()
		for i, player := range playerArr {
			playerName, _ := player.Path("name").Data().(string)
			log.Printf("player %d: %s\n", i, playerName)
		}
		c.JSON(200, v.Data())
	})

}
