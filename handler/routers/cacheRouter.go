package routers

import (
	"cacher/handler/controllers"
	"github.com/gin-gonic/gin"
)

func SetupCacheRouter(cc controllers.CacheController) *gin.Engine {

	r := gin.Default()
	rg := r.Group("/api")
	rg.GET("/cache/:key", cc.GetKey)
	rg.POST("/cache/:key", cc.SetKey)

	return r
}
