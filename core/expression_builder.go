package core

import (
	"ExpGo/core/err"
	"ExpGo/core/token"
	"ExpGo/core/tokenizer"
)

type ExpressionBuilder struct {
	Expression string
	Variables  map[string]bool
}

func NewExpressionBuilderWithExpression(expression string) *ExpressionBuilder {
	return &ExpressionBuilder{expression, make(map[string]bool, 8)}
}

func NewExpressionBuilder(expression string, variables map[string]bool) *ExpressionBuilder {
	return &ExpressionBuilder{expression, variables}
}

func (eb *ExpressionBuilder) SetVariables(variables []string) *ExpressionBuilder {
	for _, v := range variables {
		eb.Variables[v] = true
	}
	return eb
}

func (eb *ExpressionBuilder) SetVariable(variable string) *ExpressionBuilder {
	eb.Variables[variable] = true
	return eb
}

func (eb *ExpressionBuilder) BuildExpression() *Expression {
	if len(eb.Expression) == 0 {
		panic(err.ExpressionCanNotBeEmpty)
	}
	eb.Variables["pi"] = true
	eb.Variables["PI"] = true
	eb.Variables["π"] = true
	eb.Variables["e"] = true
	eb.Variables["E"] = true
	for variableName := range eb.Variables {
		if token.NewFunctionTokenWithName(variableName) != nil {
			panic(err.VariableCanNotHasSameName)
		}
	}
	return NewExpression(tokenizer.ConvertIE2IPF(eb.Expression, eb.Variables))
}
