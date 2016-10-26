package router

import (
	"github.com/gin-gonic/gin"
	"model"
	"utils/jex"
	"fmt"
)

func SetupDb(r *gin.Engine) {
	db := r.Group("/db")
	{
		db.POST("/player", playerEndpoint)
		db.GET("/player/sync", syncPlayerEndpoint)
		db.POST("/elo", eloEndpoint)
		//db.POST("/submit", submitEndpoint)
		//db.POST("/read", readEndpoint)
	}
}

func playerEndpoint(c *gin.Context) {
	var jo = jex.Load([]byte(`{"PlayerArr":` + model.Db().PlayerDb.JsonArrString() + "}"))
	c.JSON(200, jo.Data())
}
func syncPlayerEndpoint(c *gin.Context) {
	var gameIdArr = []int{23, 21, 22, 29, 39}

	var playerArr = []*jex.JsonEx{}
	for _, gameId := range gameIdArr {
		var _,playerJoArr  =GetRoundPlayerData(string(gameId))
		for _,jo :=range playerJoArr{
			playerArr = append(playerArr,jo)
		}
	}
	//http://api.liangle.com/api/passerbyking/game/players/
	//var jo = jex.Load([]byte(`{"PlayerArr":` + model.Db().PlayerDb.JsonArrString() + "}"))
	c.JSON(200, gin.H{"ok":"ok"})

}
func eloEndpoint(c *gin.Context) {
	gameIdArr := c.PostForm("gameIdArr")
	fmt.Println("gameIdArr", c.Request.Body, gameIdArr)
	var jo = jex.Load([]byte(`{"PlayerArr":` + model.Db().PlayerDb.JsonArrString() + "}"))
	c.JSON(200, jo.Data())
}