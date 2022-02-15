package input

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

	pack["command"] = parts[0]

	for index := 1; index < len(parts); index++ {
		if strings.HasPrefix(parts[index], "--") {
			pack[parts[index]] = parts[index+1]
		}
	}

	return pack
}
