package router

import (
	"github.com/gin-gonic/gin"
	"model"
)

func SetupAdmin(r *gin.Engine) {
	db := r.Group("/db")
	{
		db.POST("/player", playerEndpoint)
		//db.POST("/submit", submitEndpoint)
		//db.POST("/read", readEndpoint)
	}
}

func playerEndpoint(c *gin.Context) {
	model.Db().PlayerDb.Path()
	//model.PlayerDb.Find()
	//c.JSON(200, gin.H{
	//	"PlayerMap":  model.PlayerDb.Find(),
	//})
}