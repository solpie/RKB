package router

import (
	"github.com/gin-gonic/gin"
	"model"
	"utils/jex"
)

func SetupDb(r *gin.Engine) {
	db := r.Group("/db")
	{
		db.POST("/player", playerEndpoint)
		db.POST("/elo", eloEndpoint)
		//db.POST("/submit", submitEndpoint)
		//db.POST("/read", readEndpoint)
	}
}

func playerEndpoint(c *gin.Context) {
	var jo = jex.Load([]byte(`{"PlayerArr":` + model.Db().PlayerDb.JsonArrString() + "}"))
	c.JSON(200, jo.Data())
}
func eloEndpoint(c *gin.Context) {
	var jo = jex.Load([]byte(`{"PlayerArr":` + model.Db().PlayerDb.JsonArrString() + "}"))
	c.JSON(200, jo.Data())
}