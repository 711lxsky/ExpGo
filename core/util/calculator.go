package util

import "math/big"

func CalculateForOperator(num1, num2 string, operator string) string {
	bigNum1, _ := StringToBigFloat(num1)
	bigNum2, _ := StringToBigFloat(num2)
	switch operator {
	case "+":
		return BigFloatToString(new(big.Float).Add(bigNum1, bigNum2))
	case "-":
		return BigFloatToString(new(big.Float).Sub(bigNum1, bigNum2))
	case "*":
		return BigFloatToString(new(big.Float).Mul(bigNum1, bigNum2))
	case "/":
		return BigFloatToString(new(big.Float).Quo(bigNum1, bigNum2))
	case "^":
		bigNum3, _ := StringToBigInt(num1)
		bigNum4, _ := StringToBigInt(num2)
		return BigIntToString(new(big.Int).Exp(bigNum3, bigNum4, nil))
	case "%":
		bigNum3, _ := StringToBigInt(num1)
		bigNum4, _ := StringToBigInt(num2)
		return BigIntToString(new(big.Int).Mod(bigNum3, bigNum4))
	}
	return ""
}

func CalculateForUnary(num string, operator string) string {
	bigNum, _ := StringToBigFloat(num)
	if operator == "-" {
		return BigFloatToString(new(big.Float).Neg(bigNum))
	}
	return BigFloatToString(bigNum)
}

func CalculateForFunction(functionName string, args []string) string {
	// TODO 因为big库的匮乏，所以这部分暂时先放着，如果不使用高精度的话其实用Double很好做的
	return ""
}
