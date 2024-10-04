package util

import (
	"ExpGo/core/setting"
	"unicode"
)

// IsNumeric 判断是否为合法的数字类型字符
func IsNumeric(ch byte, lastCharIsE bool) bool {
	return unicode.IsDigit(rune(ch)) || ch == '.' || ch == 'e' || ch == 'E' ||
		(lastCharIsE && (ch == '+' || ch == '-'))
}

// IsArgumentSeparator 判断是否为函数中参数分隔符
func IsArgumentSeparator(ch byte) bool {
	return ch == ',' || ch == ';'
}

// IsOpenParenthesis 判断是否为左括号
func IsOpenParenthesis(ch byte) bool {
	return ch == '(' || ch == '[' || ch == '{'
}

// IsCloseParenthesis 判断是否为右括号
func IsCloseParenthesis(ch byte) bool {
	return ch == ')' || ch == ']' || ch == '}'
}

// IsAllowedOperatorSymbol 判断是否为合法的运算符
func IsAllowedOperatorSymbol(ch byte) bool {
	for _, symbol := range setting.AllowedOperators {
		if symbol == rune(ch) {
			return true
		}
	}
	return false
}

// IsIdentifiable 判断是否为合法的可用以命名的字符
func IsIdentifiable(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

// IsVariableOrFunctionCharacter 判断是否为合法的变量名或函数名
func IsVariableOrFunctionCharacter(ch byte) bool {
	return IsIdentifiable(ch) || unicode.IsDigit(rune(ch)) || ch == '.'
}
