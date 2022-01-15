package main

import (
	"PM/operands"
	"PM/operators"
	"log"
)

func main() {
	operations := []string{operators.Syslog, operators.Database, operators.Preprocess}

	for _, operation := range operations {
		log.Println(operation)
		operands.RunTest()
	}
}
