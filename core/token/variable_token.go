package token

import "ExpGo/core/setting"

// VariableToken 变量类型词元
type VariableToken struct {
	Name string
}

func (vt *VariableToken) Type() int {
	return setting.Variable
}
