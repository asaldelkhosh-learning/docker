package main

import (
	"PM/operands"
	"log"
)

func main() {
	operations := []string{"login", "database", "logout"}

	for _, operation := range operations {
		log.Println(operation)
		operands.RunTest()
	}
}
