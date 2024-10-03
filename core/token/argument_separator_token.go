package token

import "ExpGo/core/setting"

// ArgumentSeparatorToken 函数参数分隔符类型词元
type ArgumentSeparatorToken struct{}

func (at *ArgumentSeparatorToken) Type() int {
	return setting.ArgumentSeparator
}

func NewArgumentSeparatorToken() *ArgumentSeparatorToken {
	return &ArgumentSeparatorToken{}
}
