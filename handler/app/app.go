package app

import (
	"cacher/handler/routers"
	"cacher/internal/utils"
	"context"
	"flag"
	"fmt"
	"log"
)

type Application interface {
	Setup()
	GetContext() context.Context
}

type application struct {
	ctx    context.Context
	config *utils.ServiceConfig
}

func NewApplication(ctx context.Context) Application {
	return &application{ctx: ctx}
}

func (a *application) GetContext() context.Context {
	return a.ctx
}

func (a *application) Setup() {
	path := flag.String("config", "config.json", "env file path")
	flag.Parse()

	err := a.SetupViper(*path)
	if err != nil {
		log.Panic(err.Error())
	}

	cc := a.InitController()
	router := routers.SetupCacheRouter(cc)

	addr := fmt.Sprintf(":%d", a.config.System.HTTPPort)
	err = router.Run(addr)
	if err != nil {
		log.Panic(err.Error())
	}
}
