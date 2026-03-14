package strategies

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, handlers Handlers) {
	if router == nil || handlers == nil {
		return
	}

	v1 := router.Group("/v1")
	{
		v1.POST("/strategies", handlers.Create)
		v1.GET("/strategies/:strategy_id", handlers.GetById)
	}
}
