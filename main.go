package main

import (
	"context"
	"os"
	"os/signal"
	"project-management/internal/app"
	"syscall"
)

func main() {
	ctx := context.Background()
	serv, _ := app.New(ctx)
	defer serv.Stop()

	go func() {
		serv.Start()
	}()
	stopChan := make(chan os.Signal, 2)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	close(stopChan)
}
