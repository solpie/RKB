package router

import (
	"github.com/gin-gonic/gin"
	"model"
	"utils/jex"
	"fmt"
	"net/http"
)

func SetupDb(r *gin.Engine) {
	setupAdminCmd(r)

	db := r.Group("/db")
	{
		db.POST("/player", playerEndpoint)
		db.GET("/player/sync", syncPlayerEndpoint)
		db.POST("/elo", eloEndpoint)
		//db.POST("/submit", submitEndpoint)
		//db.POST("/read", readEndpoint)
	}

}
func setupAdminCmd(r *gin.Engine) {
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.tmpl", gin.H{
			"version": "0.4",
		})
	})
	r.POST("/admin/", func(c *gin.Context) {
		var jo = jex.Load(c)
		var cmd = jo.GetString("cmd")
		switch cmd {
		case "opUrl":
			//var opUrlArr = []string{"22", "22"}
			fmt.Println("cmd", cmd)
			var jo = jex.Load(`{"opUrlArr":["22","11"]}`)
			c.JSON(200, jo.Data())
		default:
			fmt.Println("unknown cmd", cmd)
		}
	})
}

func playerEndpoint(c *gin.Context) {
	var jo = jex.Load([]byte(`{"PlayerArr":` + model.Db().PlayerDb.JsonArrString() + "}"))
	c.JSON(200, jo.Data())
}
func syncPlayerEndpoint(c *gin.Context) {
	var gameIdArr = []string{"23", "21", "22", "29", "39", "47"}

	//var playerArr = []*jex.JsonEx{}
	//playerDoc model.PlayerDoc
	var playerDoc *model.PlayerDoc
	for _, gameId := range gameIdArr {
		var _, playerJoArr = GetRoundPlayerData(gameId)
		for _, jo := range playerJoArr {
			var playerName = jo.GetString("name")
			if _, ok := model.Db().PlayerMap[playerName]; !ok {
				model.Db().PlayerMap[playerName] = new(model.PlayerDoc).Init()
			}
			playerDoc = model.Db().PlayerMap[playerName]
			playerDoc.SetP(jo.GetString("playerNum"), "id")
			playerDoc.SetP(jo.GetNumber("playerNum"), "playerNum")
			playerDoc.SetP(jo.GetString("name"), "name")
			playerDoc.SetP(jo.GetString("ftName"), "group")
			playerDoc.SetP(jo.GetString("intro"), "intro")
			playerDoc.SetP(jo.GetString("height"), "height")
			playerDoc.SetP(jo.GetString("weight"), "weight")
			playerDoc.SetP(jo.GetString("avatar"), "avatar")
		}
	}
	//http://api.liangle.com/api/passerbyking/game/players/
	//var jo = jex.Load([]byte(`{"PlayerArr":` + model.Db().PlayerDb.JsonArrString() + "}"))
	c.JSON(200, playerDoc.Data())

}
func eloEndpoint(c *gin.Context) {
	gameIdArr := c.PostForm("gameIdArr")
	fmt.Println("gameIdArr", c.Request.Body, gameIdArr)
	var jo = jex.Load([]byte(`{"PlayerArr":` + model.Db().PlayerDb.JsonArrString() + "}"))
	c.JSON(200, jo.Data())
}