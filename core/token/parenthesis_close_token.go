package token

// ParenthesisCloseToken 闭括号（右括号）类型词元
type ParenthesisCloseToken struct{}

func (pct *ParenthesisCloseToken) Type() int {
	return ParenthesisClose
}
