package tokenizer

import (
	"ExpGo/core/err"
	"ExpGo/core/setting"
	"ExpGo/core/token"
	"ExpGo/core/util"
	"errors"
	"fmt"
	"unicode"
)

type Tokenizer struct {
	expression       []byte
	expressionLength int
	pos              int
	lastToken        token.Token
}

// 判断是否还有下一个字符
func (t *Tokenizer) hasNext() bool {
	return t.pos < t.expressionLength
}

func (t *Tokenizer) isEndOfExpression(index int) bool {
	return index >= t.expressionLength
}

// 解析下一个词元
func (t *Tokenizer) nextToken() (token.Token, error) {
	// 定位当前位置字符
	curChar := t.expression[t.pos]
	for unicode.IsSpace(rune(curChar)) {
		// 跳过空白字符
		t.pos++
		curChar = t.expression[t.pos]
	}
	// 判断是否为数字或者小数点
	if unicode.IsDigit(rune(curChar)) || curChar == '.' {
		if t.lastToken != nil {
			// 上一个词元也是数字， 非法
			if t.lastToken.Type() == setting.Number {
				return nil, errors.New(fmt.Sprintf(err.CanNotParseToken, curChar, t.pos))
			} else if t.isNeedImplicitMultiplication() {
				// 需要插入隐式乘法
				t.lastToken = token.NewMultiplicationOperatorToken()
				return t.lastToken, nil
			}
		}
		return t.parseNumberToken(curChar)
	} else if util.IsArgumentSeparator(curChar) {
		// 分隔符
		return t.parseArgumentSeparatorToken()
	} else if util.IsOpenParenthesis(curChar) {
		// 左括号, 先检查是否需要插入隐式乘法
		if t.isNeedImplicitMultiplication() {
			// 需要插入隐式乘法
			t.lastToken = token.NewMultiplicationOperatorToken()
			return t.lastToken, nil
		}
		// 解析开括号词元
		return t.parseParenthesisOpenToken(true)
	} else if util.IsCloseParenthesis(curChar) {
		// 解析闭括号词元
		return t.parseParenthesisOpenToken(false)
	} else if util.IsAllowedOperatorSymbol(curChar) {
		return t.parseOperatorToken(curChar)
	}
}

// 解析数字词元
func (t *Tokenizer) parseNumberToken(firstChar byte) (token.Token, error) {
	// 先定位到当前位置的字符以及接下来需要解析的长度
	offset, length := t.pos, 1
	// 然后将解析位置递增1
	t.pos += 1
	// 判断是否到达表达式末尾
	if t.isEndOfExpression(offset + length) {
		// 到达，直接返回
		t.lastToken = token.NewNumberToken(string(firstChar))
		return t.lastToken, nil
	}
	// 遍历拿到数字字符
	for !t.isEndOfExpression(offset+length) &&
		util.IsNumeric(t.expression[offset+length],
			t.expression[offset+length-1] == 'e' || t.expression[offset+length-1] == 'E') {
		length++
		t.pos++
	}
	// 检查字符是否以 e 或者 E 结尾，如果是则将位置指针回退一位
	if t.expression[offset+length-1] == 'e' || t.expression[offset+length-1] == 'E' {
		length--
		t.pos--
	}
	// 构建数字词元返回
	t.lastToken = token.NewNumberToken(string(t.expression[offset : offset+length]))
	return t.lastToken, nil
}

// 判断是否需要隐式乘法
func (t *Tokenizer) isNeedImplicitMultiplication() bool {
	return t.lastToken.Type() != setting.Operator &&
		t.lastToken.Type() != setting.ParenthesisOpen &&
		t.lastToken.Type() != setting.Function &&
		t.lastToken.Type() != setting.ArgumentSeparator
}

func (t *Tokenizer) parseArgumentSeparatorToken() (token.Token, error) {
	// 位置递增
	t.pos++
	t.lastToken = token.NewArgumentSeparatorToken()
	return t.lastToken, nil
}

func (t *Tokenizer) parseParenthesisOpenToken(isOpen bool) (token.Token, error) {
	if isOpen {
		t.lastToken = token.NewParenthesisOpenToken()
	} else {
		t.lastToken = token.NewParenthesisCloseToken()
	}
	t.pos++
	return t.lastToken, nil
}

func (t *Tokenizer) parseOperatorToken(ch byte) (token.Token, error) {
	// TODO 这里后续可以设计以实现对用户自定义的操作符支持
	t.pos++
	t.lastToken =

}
