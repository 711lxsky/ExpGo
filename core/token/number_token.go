package token

// NumberToken 数字常量类型词元
type NumberToken struct {
	Value string
}

func (nt *NumberToken) Type() int {
	return Number
}
