package token

// OperatorToken 运算符类型词元
type OperatorToken struct {
	symbol          string
	operationNumber int
	priority        int
}

func (ot *OperatorToken) Type() int {
	return Operator
}

func BuildOperatorToken(symbol string, operatorNumber int, priority int) *OperatorToken {
	return &OperatorToken{
		symbol:          symbol,
		operationNumber: operatorNumber,
		priority:        priority,
	}
}
