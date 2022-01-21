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
	l.Operand = map[string]func(){
		operators.Syslog:     operands.Log,
		operators.Database:   operands.DBTransaction,
		operators.Preprocess: operands.UserProcess,
	}
}
