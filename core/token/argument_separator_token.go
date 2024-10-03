package token

// ArgumentSeparatorToken 函数参数分隔符类型词元
type ArgumentSeparatorToken struct{}

func (at *ArgumentSeparatorToken) Type() int {
	return ArgumentSeparator
}
