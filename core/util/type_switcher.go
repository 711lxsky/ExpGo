package util

import "math/big"

func Float64ToBigFloat(f float64) *big.Float {
	return big.NewFloat(f)
}

func BigFloatToString(bf *big.Float) string {
	return bf.Text('f', -1)
}

func BigIntToString(bi *big.Int) string {
	return bi.String()
}

func StringToBigInt(s string) (*big.Int, bool) {
	return new(big.Int).SetString(s, 10)
}

func StringToBigFloat(s string) (*big.Float, bool) {
	return new(big.Float).SetString(s)
}
