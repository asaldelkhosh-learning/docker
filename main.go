package main

import (
	"PM/linker"
	"PM/operators"
	"log"
	"time"
)

func main() {
	operations := []string{operators.Syslog, operators.Database, operators.Preprocess}

	for _, operation := range operations {
		log.Println(operation)
		operator := linker.Link(operation)
		go operator()
	}
	time.Sleep(10 * time.Second)
}
