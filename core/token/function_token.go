package token

import "ExpGo/core/setting"

// FunctionToken 函数类型词元
type FunctionToken struct {
	Name           string
	ArgumentNumber int
}

func (ft FunctionToken) Type() int {
	return setting.Function
}

func NewFunctionToken(name string, argumentNumber int) *FunctionToken {
	return &FunctionToken{
		Name:           name,
		ArgumentNumber: argumentNumber,
	}
}

func NewFunctionTokenWithName(name string) *FunctionToken {
	if name == setting.Pow {
		return NewFunctionToken(name, 2)
	} else {
		for _, function := range setting.AllowedFunctions {
			if name == function {
				return NewFunctionToken(name, 1)
			}
		}
	}
	return nil
}
