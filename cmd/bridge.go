package main

import (
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
	var serviceID telegram.String
	if err := ctx.Argument("service", &serviceID); err != nil {
		return err
	}

	var taskID telegram.String
	if err := ctx.Argument("task", &taskID); err != nil {
		return err
	}

	doneChan, errChan := ctx.QueueTask(string(serviceID), string(taskID))
	select {
	case <-doneChan:
		ctx.Reply("DONE")
	case err := <-errChan:
		return err
	}
	return nil
}
