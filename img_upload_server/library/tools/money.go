package tools

import "github.com/shopspring/decimal"

func MoenyFormat(f float64)float64{
	v1, _ := decimal.NewFromFloat(f).Round(2).Float64()
	return v1
}
