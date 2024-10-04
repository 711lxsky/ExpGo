package token

import "ExpGo/core/setting"

// NumberToken 数字常量类型词元
type NumberToken struct {
	Value string
}

func (nt NumberToken) Type() int {
	return setting.Number
}

func NewNumberToken(value string) *NumberToken {
	return &NumberToken{Value: value}
}
