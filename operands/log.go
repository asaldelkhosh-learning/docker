package operands

import "PM/internal/commandor"

var file = "log.sh"

func Log() {
	commandor.Exec(file)
}
