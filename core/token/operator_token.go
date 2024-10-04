package token

import "ExpGo/core/setting"

// OperatorToken 运算符类型词元
type OperatorToken struct {
	Symbol            string
	OperationNumber   int
	Priority          int
	IsLeftAssociative bool
}

func (ot OperatorToken) Type() int {
	return setting.Operator
}

func NewOperatorToken(symbol string, operatorNumber int, priority int, isLeftAssociative bool) *OperatorToken {
	return &OperatorToken{
		Symbol:            symbol,
		OperationNumber:   operatorNumber,
		Priority:          priority,
		IsLeftAssociative: isLeftAssociative,
	}
}

func NewMultiplicationOperatorToken() *OperatorToken {
	return NewOperatorToken("*", 2, setting.MulOrDivOrModPriority, true)
}

func NewOperatorTokenWithSymbolAndArgs(symbol string, operatorNumber int) *OperatorToken {
	if operatorNumber == 1 {
		return NewOperatorToken(symbol, 1, setting.UnaryOrPowerPriority, false)
	} else {
		switch symbol {
		case "+":
		case "-":
			return NewOperatorToken(symbol, 2, setting.AddOrSubPriority, true)
		case "*":
		case "/":
		case "%":
			return NewOperatorToken(symbol, 2, setting.MulOrDivOrModPriority, true)
		case "^":
			return NewOperatorToken(symbol, 2, setting.UnaryOrPowerPriority, true)
		}
	}
	return nil
}
