package main

import (
	format "github.com/VitalySR/task13_mod"
	"log"
)

func main() {
	err := format.Do("patients", "result")
	if err != nil {
		log.Fatal(err)
	}
}
