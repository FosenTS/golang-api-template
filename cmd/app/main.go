package main

import (
	"context"
	"flag"
	"fmt"
	"golang-api-template/internal/application"
	"golang-api-template/internal/application/config"
	"log"
	"time"
)

var modeFlag *string

func main() {
	time.Local = time.UTC
	modeFlag = flag.String("mode", config.Mode.Local(), "specifies whether program is for dev or prod or local usage, default is local")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()
	mode := *modeFlag
	switch mode {
	case config.Mode.Local():
		fmt.Println("Local mode!")
	case config.Mode.Deploy():
		fmt.Println("Deploy mode!")
	default:
		log.Fatalln("invalid mode")
	}
	config.LoadEnv(mode)

	app, err := application.NewApp(ctx)
	if err != nil {
		return
	}

	err = app.Run(ctx)
	if err != nil {
		return
	}
}
