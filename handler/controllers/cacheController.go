package controllers

import (
	"GorillaCacher/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CacheController interface {
	GetKey(ctx *gin.Context)
	SetKey(ctx *gin.Context)
}

type cacheController struct {
	cacheService services.CacheService
}

func NewCacheController(cs services.CacheService) CacheController {
	return &cacheController{cacheService: cs}
}

func (c *cacheController) GetKey(ctx *gin.Context) {
	k := ctx.Param("key")
	v, err := c.cacheService.Get(k)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch key"})
		return
	}
	ctx.JSON(http.StatusOK, v)
}

func (c *cacheController) SetKey(ctx *gin.Context) {
	k := ctx.Param("key")
	v := ctx.PostForm("value")
	ttlStr := ctx.PostForm("ttl")
	ttl, _ := time.ParseDuration(fmt.Sprintf("%ss", ttlStr))
	err := c.cacheService.Set(k, v, ttl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set key"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"key": k, "value": v, "ttl": ttl})
}
