package main

import (
	"GorillaCacher/handler/app"
	"context"
)

func main() {
	application := app.NewApplication(context.Background())
	application.Setup()
}
