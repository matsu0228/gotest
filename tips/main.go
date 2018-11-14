package main

import (
	"fmt"
	"math"
)

// roundFloat :Floatを四捨五入して返却
func roundFloat(f float64) float64 {
	return math.Floor(f + .5)
}

type tax struct {
	taxPercent float64
}

func newTax() tax {
	// 税率変更時はここを直す
	return tax{taxPercent: 0.08}
}

// calcurateTaxExcludeAmount :税込額から、税抜額(本体価格)を算出
func (t tax) calcurateTaxExcludeAmount(taxIncludedPrice int) int {

	// taxRate := float64(t.taxPercent) // 割算をしたときに結果をfloatで返却するため
	price := float64(taxIncludedPrice)
	return int(roundFloat(price / (1 + t.taxPercent)))
}

// calcurateTaxAmount :消費税額を算出
func (t tax) calcurateTaxAmount(taxIncludedPrice int) int {
	taxExcludePrice := t.calcurateTaxExcludeAmount(taxIncludedPrice)
	return taxIncludedPrice - taxExcludePrice
}
func main() {
	tax := newTax()
	fmt.Printf("税込額: %v, 税抜額: %v \n", 2000, tax.calcurateTaxExcludeAmount(2000))
}
