package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	reader *bufio.Reader
}

func (i Input) Init() Input {
	i.reader = bufio.NewReader(os.Stdin)
	if i.reader == nil {
		panic(fmt.Errorf("problem in reader"))
	}
	return i
}

func (i Input) Get() string {
	cmd, _ := i.reader.ReadString('\n')
	cmd = strings.Trim(cmd, "\n")

	return cmd
}

func (i Input) Decode(cmd string) map[string]string {
	pack := make(map[string]string)
	parts := strings.Split(cmd, " ")

	for index, part := range parts {
		pack[strconv.Itoa(index)] = part
	}

	return pack
}
