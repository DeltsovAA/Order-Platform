package main

import (
	"order-platform/internal/consumer"
)

func main() {
	app := consumer.New()
	app.Run()
}
