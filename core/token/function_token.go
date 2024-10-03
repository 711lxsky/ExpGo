package token

// FunctionToken 函数类型词元
type FunctionToken struct{}

func (ft *FunctionToken) Type() int {
	return Function
}
