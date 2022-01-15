package linker

import (
	"PM/operands"
	"PM/operators"
)

func Link(operator string) func() {
	switch operator {
	case operators.Syslog:
		return operands.RunTest
	case operators.Database:
		return operands.RunTest
	case operators.Preprocess:
		return operands.RunTest
	}

	return nil
}
