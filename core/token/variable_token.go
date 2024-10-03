package token

// VariableToken 变量类型词元
type VariableToken struct {
	Name string
}

func (vt *VariableToken) Type() int {
	return Variable
}
