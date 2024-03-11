package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alexthvest/homelab-bridge/pkg/homelab/telegram"
)

func setupBridge() telegram.Bridge {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatalf("BOT_TOKEN is missing")
	}

	router := setupRouter()

	bridge, err := telegram.NewBridge(token, router)
	if err != nil {
		log.Fatalf("error while established bridge: %v", err)
	}

	return bridge
}

func setupRouter() *telegram.Router {
	router := telegram.NewRouter()

	svcCommand := router.Command("svc")
	svcCommand.Argument("service")

	taskCommand := svcCommand.Command("task")

	runCommand := taskCommand.Command("run")
	runCommand.Argument("task")
	runCommand.Handler(runHandler)

	return router
}

func runHandler(ctx telegram.Context) error {
	doneChan, errChan := ctx.QueueTask("stable-diffusion", "txt2img")
	go func() {
		select {
		case <-doneChan:
			ctx.Reply("DONE")
		case err := <-errChan:
			ctx.Reply(fmt.Sprintf("error: %v", err))
		}
	}()
	return nil
}
