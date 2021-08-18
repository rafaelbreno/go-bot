package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Loading .env file
	godotenv.Load()

	<-stop
}
