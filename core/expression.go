package core

import (
	"ExpGo/core/err"
	"ExpGo/core/setting"
	"ExpGo/core/token"
	"ExpGo/core/util"
	"fmt"
)

type Expression struct {
	tokens    []token.Token
	variables map[string]string
}

func SetDefaultVariables() map[string]string {
	variables := make(map[string]string, 8)
	variables["pi"] = "3.1415926"
	variables["PI"] = "3.1415926"
	variables["π"] = "3.1415926"
	variables["φ"] = "1.61803398874"
	variables["e"] = "2.7182818"
	return variables
}

func NewExpression(tokens []token.Token) *Expression {
	return &Expression{
		tokens:    tokens,
		variables: SetDefaultVariables(),
	}
}

func (e *Expression) Validate() bool {
	// 检查参数能否对应上
	count := 0
	for _, t := range e.tokens {
		switch t.Type() {
		case setting.Variable:
			name := t.(token.VariableToken).Name
			if _, ok := e.variables[name]; !ok {
				panic(fmt.Sprintf(err.VariableMissValue, name))
			}
		case setting.Number:
			count += 1
			break
		case setting.Function:
			ft := t.(token.FunctionToken)
			argsNum := ft.ArgumentNumber
			if argsNum > count {
				panic(fmt.Sprintf(err.HaveNotEnoughArguments, ft.Name))
			}
			if argsNum > 1 {
				count -= argsNum - 1
			} else if argsNum == 0 {
				count++
			}
			break
		case setting.Operator:
			ot := t.(token.OperatorToken)
			if ot.OperationNumber == 2 {
				count -= 1
			}
			break
		}
		if count < 1 {
			panic(err.TooManyOperators)
		}
	}
	if count > 1 {
		panic(err.TooManyArguments)
	}
	return true
}

func (e *Expression) evaluate() string {
	output := util.NewStack[string]()
	for _, t := range e.tokens {
		if t.Type() == setting.Number {
			// 数字
			output.Push(t.(token.NumberToken).Value)
		} else if t.Type() == setting.Variable {
			// 变量
			name := t.(token.VariableToken).Name
			value := e.variables[name]
			if value == "" {
				panic(fmt.Sprintf(err.VariableMissValue, name))
			}
			output.Push(value)
		} else if t.Type() == setting.Operator {
			// 操作运算符
			ot := t.(token.OperatorToken)
			if output.Size() < ot.OperationNumber {
				panic(fmt.Sprintf(err.LackSufficientOperands, ot.Symbol))
			}
			if ot.OperationNumber == 2 {
				rightArg := output.Pop()
				leftArg := output.Pop()
				output.Push(util.CalculateForOperator(leftArg, rightArg, ot.Symbol))
			} else if ot.OperationNumber == 1 {
				arg := output.Pop()
				output.Push(util.CalculateForUnary(arg, ot.Symbol))
			}
		} else if t.Type() == setting.Function {
			// 函数
			ft := t.(token.FunctionToken)
			argNums := ft.ArgumentNumber
			if output.Size() < argNums {
				panic(fmt.Sprintf(err.LackSufficientOperands, ft.Name))
			}
			args := make([]string, argNums)
			for i := 0; i < argNums; i++ {
				args[i] = output.Pop()
			}
			output.Push(util.CalculateForFunction(ft.Name, args))
		}
	}
	if output.Size() >= 1 {
		panic(err.InvalidItemsInStack)
	}
	return output.Pop()
}
