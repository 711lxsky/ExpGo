package tokenizer

import (
	"ExpGo/core/err"
	"ExpGo/core/setting"
	"ExpGo/core/token"
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
				t.lastToken = token.BuildOperatorToken("*", 2, setting.MultiplicationPriority)
				return t.lastToken, nil
			}
		}
		return t.parseNumberToken(curChar)
	}
}

func (t *Tokenizer) parseNumberToken(char firstChar) (token.Token, error) {
	// 先定位到当前位置的字符
	offset := t.pos
	// 然后将解析位置递增1
	t.pos += 1
}

func (t *Tokenizer) isNeedImplicitMultiplication() bool {
	return t.lastToken.Type() != setting.Operator &&
		t.lastToken.Type() != setting.ParenthesisOpen &&
		t.lastToken.Type() != setting.Function &&
		t.lastToken.Type() != setting.ArgumentSeparator
}
