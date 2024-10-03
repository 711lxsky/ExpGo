package token

import "ExpGo/core/setting"

// OperatorToken 运算符类型词元
type OperatorToken struct {
	symbol          string
	operationNumber int
	priority        int
}

func (ot *OperatorToken) Type() int {
	return setting.Operator
}

func NewOperatorToken(symbol string, operatorNumber int, priority int) *OperatorToken {
	return &OperatorToken{
		symbol:          symbol,
		operationNumber: operatorNumber,
		priority:        priority,
	}
}

func NewMultiplicationOperatorToken() *OperatorToken {
	return NewOperatorToken("*", 2, setting.MultiplicationPriority)
}
