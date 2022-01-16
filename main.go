package main

import (
	"PM/linker"
	"PM/operators"
	"log"
	"time"
)

func main() {
	operations := []string{operators.Syslog, operators.Database, operators.Preprocess}

	var link linker.Linker
	link.Config()

	for _, operation := range operations {
		log.Println(operation)
		operator := link.Link(operation)
		go operator()
	}
	time.Sleep(10 * time.Second)
}
