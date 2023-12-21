package main

import (
	"log"

	"github.com/boltdbgui/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
