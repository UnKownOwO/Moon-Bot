package main

import (
	"moon-bot/bs/app"
	"moon-bot/pkg/statsviz"
)

func main() {
	go func() {
		_ = statsviz.Run("0.0.0.0:6662")
	}()
	app.Run()
}
