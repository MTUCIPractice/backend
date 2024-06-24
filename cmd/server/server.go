package main

import (
	"context"
	"github.com/practice/backend/intertnal/controller/http"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	var (
		err    error
		ctx    context.Context
		cancel context.CancelFunc
	)

	ctx, cancel = signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	defer cancel()

	//initialize server
	server, err := http.New()
	if err != nil {
		log.Fatal(err)
	}
}
