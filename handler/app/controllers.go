package app

import (
	"cacher/handler/controllers"
	"cacher/internal/services"
)

func (a *application) InitController() controllers.CacheController {
	hr := a.InitCacheService()
	return controllers.NewCacheController(services.NewCacheService(hr))
}
