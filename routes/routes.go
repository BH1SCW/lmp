package routes

import (
	"github.com/gin-gonic/gin"
	"lmp/controllers"
	"net/http"

	"lmp/logger"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msh": "404",
		})
	})

	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)
	r.POST("/uploadfiles", controllers.UpLoadFiles)
	r.GET("/allplugins", controllers.PrintAllplugins)
	r.POST("/data/collect", controllers.Collect)

	return r
}
