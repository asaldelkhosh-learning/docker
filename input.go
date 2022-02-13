package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Input struct {
	reader *bufio.Reader
}

func (i Input) Init() Input {
	reader := bufio.NewReader(os.Stdin)
	if reader == nil {
		panic(fmt.Errorf("problem in reader"))
	}
	return i
}

func (i Input) Get() string {
	cmd, _ := i.reader.ReadString('\n')
	cmd = strings.Trim(cmd, "\n")

	return cmd
}
