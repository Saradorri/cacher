package app

import (
	"GorillaCacher/handler/controllers"
	"GorillaCacher/internal/services"
)

func (a *application) InitController() controllers.CacheController {
	hr := a.InitCacheService()
	return controllers.NewCacheController(services.NewCacheService(hr))
}
