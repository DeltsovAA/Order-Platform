package main

import (
	"log"
	"order-platform/internal/producer"
)

func main() {
	if err := producer.New().Run(); err != nil {
		log.Println(err)
		return
	}
}
