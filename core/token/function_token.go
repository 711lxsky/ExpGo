package token

import "ExpGo/core/setting"

// FunctionToken 函数类型词元
type FunctionToken struct{}

func (ft *FunctionToken) Type() int {
	return setting.Function
}
