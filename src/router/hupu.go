package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"log"
	"utils/jex"
	"fmt"
)

func SetupHupuAPI(r *gin.Engine) {
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
	/*
 23 hefei1
 21 sh6
 22 sh7
 29 hz1
 39 sh8
 * */
	r.GET("/api/round/:round", func(c *gin.Context) {
		round := c.Param("round")
		response, _ := http.Get("http://api.liangle.com/api/passerbyking/game/match/" + round)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		c.JSON(200, jex.Load(body).Data())
	})
	r.GET("/api/passerbyking/*api", func(c *gin.Context) {
		api := c.Param("api")
		fmt.Println(api)
		response, _ := http.Get("http://api.liangle.com/api/passerbyking" + api)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		c.JSON(200, jex.Load(body).Data())
	})

}
