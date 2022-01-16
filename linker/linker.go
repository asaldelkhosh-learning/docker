package linker

import (
	"PM/operands"
	"PM/operators"
)

func Link(operator string) func() {
	switch operator {
	case operators.Syslog:
		return operands.Log
	case operators.Database:
		return operands.DBTransaction
	case operators.Preprocess:
		return operands.UserProcess
	}

	return nil
}
