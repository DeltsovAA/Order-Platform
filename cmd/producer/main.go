package main

import (
	"log"
	"order-platform/internal/producer"
)

func main() {
	p := producer.New()

	if err := p.Run(); err != nil {
		log.Println(err)
		return
	}
}
