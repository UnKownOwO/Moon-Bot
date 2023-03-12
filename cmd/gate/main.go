package main

import (
	"context"
	"fmt"
	"moon-bot/gate/app"
	"moon-bot/pkg/statsviz"
	"os"
)

func main() {
	go func() {
		_ = statsviz.Run("0.0.0.0:6661")
	}()
	err := app.Run(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
