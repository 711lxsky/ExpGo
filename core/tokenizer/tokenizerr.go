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
	variables        map[string]bool
}

func NewTokenizer(expression string, variables map[string]bool) *Tokenizer {
	return &Tokenizer{
		expression:       []byte(expression),
		expressionLength: len(expression),
		pos:              0,
		lastToken:        nil,
		variables:        variables,
	}
}

// 判断是否还有下一个字符
func (t *Tokenizer) hasNext() bool {
	return t.pos < t.expressionLength
}

func (t *Tokenizer) isEndOfExpression(index int) bool {
	return index >= t.expressionLength
}

// 拿到下一个词元
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
		// 解析运算操作符
		return t.parseOperatorToken(curChar)
	} else if util.IsIdentifiable(curChar) {
		// 先判断是否需要插入隐式乘法
		if t.isNeedImplicitMultiplication() {
			// 需要插入隐式乘法
			t.lastToken = token.NewMultiplicationOperatorToken()
			return t.lastToken, nil
		}
		// 解析变量或者函数词元
		return t.parseFunctionOrVariableToken()
	}
	// 没有相匹配的词元
	return nil, errors.New(fmt.Sprintf(err.CanNotParseToken, curChar, t.pos))
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
	// 先假设操作数为2
	args := 2
	if t.lastToken == nil {
		// 如果上一个次元为空则改为1
		args = 1
	} else {
		// 如果上一个词元为括号或者分隔符，则改为1
		if t.lastToken.Type() == setting.ParenthesisOpen || t.lastToken.Type() == setting.ArgumentSeparator {
			args = 1
		} else if t.lastToken.Type() == setting.Operator {
			// 上一个词元为二元或者一元且左结合
			lastOp := t.lastToken.(token.OperatorToken)
			if lastOp.OperationNumber == 2 || (lastOp.OperationNumber == 1 && lastOp.IsLeftAssociative) {
				args = 1
			}
		}
	}
	// 构建返回
	t.lastToken = token.NewOperatorTokenWithSymbolAndArgs(string(ch), args)
	return t.lastToken, nil
}

func (t *Tokenizer) parseFunctionOrVariableToken() (token.Token, error) {
	// 初始解析位置、长度
	offset, length := t.pos, 1
	if t.isEndOfExpression(offset) {
		t.pos++
	}
	// 临时变量存储最近有效词元
	var lastValidToken token.Token
	lastValidToken, lastValidTokenLen, ptr := nil, 1, offset+length-1
	// 迭代处理
	for !t.isEndOfExpression(ptr) && util.IsVariableOrFunctionCharacter(t.expression[ptr]) {
		// 分片截取
		name := string(t.expression[offset : length+1])
		if t.variables != nil && t.variables[name] {
			// 解析为变量词元
			lastValidTokenLen = length
			lastValidToken = token.NewVariableToken(name)
		} else {
			// 尝试解析为函数词元
			lastValidToken = token.NewFunctionTokenWithName(name)
			if lastValidToken != nil {
				lastValidTokenLen = length
			}
		}
		// 长度加一继续遍历
		length++
		ptr = offset + length - 1
	}
	// 没有解析出，说明命名有问题
	if lastValidToken == nil {
		return nil, errors.New(fmt.Sprintf(err.CheckTheName, t.expression[offset], offset))
	}
	// 解析成功，加上偏移量
	t.pos += lastValidTokenLen
	t.lastToken = lastValidToken
	return t.lastToken, nil
}
