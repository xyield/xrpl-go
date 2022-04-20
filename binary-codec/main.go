package main

import (
	"log"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
)

func main() {
	err := definitions.LoadDefinitions()
	if err != nil {
		log.Println(err)
	}
}
