package linker

import (
	"PM/operands"
	"PM/operators"
)

type Linker struct {
	Operand map[string]func()
}

func (l *Linker) Link(operator string) func() {
	return l.Operand[operator]
}

func (l *Linker) Config() {
	for _, operator := range operators.List() {
		var operand func()

		switch operator {
		case operators.Syslog:
			operand = operands.Log
		case operators.Database:
			operand = operands.DBTransaction
		case operators.Preprocess:
			operand = operands.UserProcess
		}

		l.Operand[operator] = operand
	}
}
