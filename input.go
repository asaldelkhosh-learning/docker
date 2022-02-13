package main

import (
	"bufio"
	"fmt"
	"os"
)

type Input struct {
	reader *bufio.Reader
}

func (i *Input) Init() Input {
	reader := bufio.NewReader(os.Stdin)
	if reader == nil {
		panic(fmt.Errorf("problem in reader"))
	}
	return *i
}
