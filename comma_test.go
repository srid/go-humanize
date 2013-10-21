package humanize

import (
	"math/big"
	"testing"
)

func TestCommas(t *testing.T) {
	testList{
		{"0", Comma(0), "0"},
		{"10", Comma(10), "10"},
		{"100", Comma(100), "100"},
		{"1,000", Comma(1000), "1,000"},
		{"10,000", Comma(10000), "10,000"},
		{"100,000", Comma(100000), "100,000"},
		{"10,000,000", Comma(10000000), "10,000,000"},
		{"10,100,000", Comma(10100000), "10,100,000"},
		{"10,010,000", Comma(10010000), "10,010,000"},
		{"10,001,000", Comma(10001000), "10,001,000"},
		{"123,456,789", Comma(123456789), "123,456,789"},
		{"maxint", Comma(9.223372e+18), "9,223,372,000,000,000,000"},
		{"minint", Comma(-9.223372e+18), "-9,223,372,000,000,000,000"},
		{"-123,456,789", Comma(-123456789), "-123,456,789"},
		{"-10,100,000", Comma(-10100000), "-10,100,000"},
		{"-10,010,000", Comma(-10010000), "-10,010,000"},
		{"-10,001,000", Comma(-10001000), "-10,001,000"},
		{"-10,000,000", Comma(-10000000), "-10,000,000"},
		{"-100,000", Comma(-100000), "-100,000"},
		{"-10,000", Comma(-10000), "-10,000"},
		{"-1,000", Comma(-1000), "-1,000"},
		{"-100", Comma(-100), "-100"},
		{"-10", Comma(-10), "-10"},
	}.validate(t)
}

func BenchmarkCommas(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comma(1000000000)
		Comma(1234567890)
	}
}

func bigComma(i int64) string {
	return BigComma(big.NewInt(i))
}

func TestBigCommas(t *testing.T) {
	testList{
		{"0", bigComma(0), "0"},
		{"10", bigComma(10), "10"},
		{"100", bigComma(100), "100"},
		{"1,000", bigComma(1000), "1,000"},
		{"10,000", bigComma(10000), "10,000"},
		{"100,000", bigComma(100000), "100,000"},
		{"10,000,000", bigComma(10000000), "10,000,000"},
		{"10,100,000", bigComma(10100000), "10,100,000"},
		{"10,010,000", bigComma(10010000), "10,010,000"},
		{"10,001,000", bigComma(10001000), "10,001,000"},
		{"123,456,789", bigComma(123456789), "123,456,789"},
		{"maxint", bigComma(9.223372e+18), "9,223,372,000,000,000,000"},
		{"minint", bigComma(-9.223372e+18), "-9,223,372,000,000,000,000"},
		{"-123,456,789", bigComma(-123456789), "-123,456,789"},
		{"-10,100,000", bigComma(-10100000), "-10,100,000"},
		{"-10,010,000", bigComma(-10010000), "-10,010,000"},
		{"-10,001,000", bigComma(-10001000), "-10,001,000"},
		{"-10,000,000", bigComma(-10000000), "-10,000,000"},
		{"-100,000", bigComma(-100000), "-100,000"},
		{"-10,000", bigComma(-10000), "-10,000"},
		{"-1,000", bigComma(-1000), "-1,000"},
		{"-100", bigComma(-100), "-100"},
		{"-10", bigComma(-10), "-10"},
	}.validate(t)
}

func TestVeryBigCommas(t *testing.T) {
	tests := []struct{ in, exp string }{
		{
			"84889279597249724975972597249849757294578485",
			"84,889,279,597,249,724,975,972,597,249,849,757,294,578,485",
		},
		{
			"-84889279597249724975972597249849757294578485",
			"-84,889,279,597,249,724,975,972,597,249,849,757,294,578,485",
		},
	}
	for _, test := range tests {
		n, _ := (&big.Int{}).SetString(test.in, 10)
		got := BigComma(n)
		if test.exp != got {
			t.Errorf("Expected %q, got %q", test.exp, got)
		}
	}
}
