package main

import (
	"moon-bot/gate/app"
	"moon-bot/pkg/statsviz"
)

func main() {
	go func() {
		_ = statsviz.Run("0.0.0.0:6661")
	}()
	app.Run()
}
