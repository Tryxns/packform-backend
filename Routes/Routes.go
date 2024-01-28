package Routes

import (
	_ "net/http"
	"packform/Controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
	}))
	grp := router.Group("/orders")
	{
		grp.GET("", Controllers.GetOrders)
	}

	return router
}
