package token

import "ExpGo/core/setting"

// ParenthesisOpenToken 开括号（左括号）类型词元
type ParenthesisOpenToken struct{}

func (pot *ParenthesisOpenToken) Type() int {
	return setting.ParenthesisOpen
}

func NewParenthesisOpenToken() *ParenthesisOpenToken {
	return &ParenthesisOpenToken{}
}
