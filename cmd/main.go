package main

import (
	"cacher/handler/app"
	"context"
)

func main() {
	application := app.NewApplication(context.Background())
	application.Setup()
}
