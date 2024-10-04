package tokenizer

import (
	err2 "ExpGo/core/err"
	"ExpGo/core/setting"
	"ExpGo/core/token"
	"ExpGo/core/util"
	"errors"
	"fmt"
)

func ConvertIE2IPF(expression string, variables map[string]bool) (res []token.Token) {
	var (
		stack     *util.Stack[token.Token]
		output    []token.Token
		tokenizer *Tokenizer
	)
	// 初始化工具
	stack, output, tokenizer = util.NewStack[token.Token](), make([]token.Token, 0), NewTokenizer(expression, variables)
	// 迭代解析
	for tokenizer.hasNext() {
		curToken, err := tokenizer.nextToken()
		if err != nil {
			panic(err)
		}
		switch curToken.Type() {
		case setting.Number:
		case setting.Variable:
			// 数字和变量直接放入输出列表
			output = append(output, curToken)
			break
		case setting.Function:
			// 函数放入栈中
			stack.Push(curToken)
			break
		case setting.ArgumentSeparator:
			// 函数参数分隔符
			for !stack.IsEmpty() && stack.Peek().Type() != setting.ParenthesisOpen {
				output = append(output, stack.Pop())
			}
			// 栈为空或栈顶元素不是左括号
			if stack.IsEmpty() || stack.Peek().Type() != setting.ParenthesisOpen {
				panic(errors.New(fmt.Sprintf(err2.MissSeparatorOrParenthesis, tokenizer.pos)))
			}
			break
		case setting.Operator:
			// 解析为运算符
			for !stack.IsEmpty() && stack.Peek().Type() == setting.Operator {
				// 栈顶元素同样为运算符
				op1, op2 := curToken.(token.OperatorToken), stack.Peek().(token.OperatorToken)
				if op1.OperationNumber == 1 && op2.OperationNumber == 2 {
					break
				} else if (op1.IsLeftAssociative && op1.Priority <= op2.Priority) || (op1.Priority < op2.Priority) {
					output = append(output, stack.Pop())
				} else {
					break
				}
			}
		case setting.ParenthesisOpen:
			// 左括号
			stack.Push(curToken)
			break
		case setting.ParenthesisClose:
			// 右括号
			for stack.Peek().Type() != setting.ParenthesisOpen {
				output = append(output, stack.Pop())
			}
			stack.Pop()
			if !stack.IsEmpty() && stack.Peek().Type() == setting.Function {
				output = append(output, stack.Pop())
			}
			break
		default:
			panic(errors.New(fmt.Sprintf(err2.UnknownTokenType, tokenizer.pos)))
		}
	}
	// 最后将栈中剩余元素全部弹出
	for !stack.IsEmpty() {
		curToken := stack.Pop()
		if curToken.Type() == setting.ParenthesisClose || curToken.Type() == setting.ParenthesisOpen {
			panic(errors.New(fmt.Sprintf(err2.MissSeparatorOrParenthesis, tokenizer.pos)))
		} else {
			output = append(output, curToken)
		}
	}
	return output
}
