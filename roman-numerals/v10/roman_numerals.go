package v1

import "strings"

type romanNumeral struct {
	Value  int
	Symbol string
}

type romanNumerals []romanNumeral

var allRomanNumerals = romanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

//converts from arabic to roman
func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

// ConvertToArabic converts a Roman Numeral to an Arabic number.
func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		// look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
		if couldBeSubtractive(i, symbol, roman) {

			// get the value of the two character string
			value := allRomanNumerals.ValueOf(symbol, roman[i+1])

			if value != 0 {
				total += value
				i++ // move past this character too for the next loop
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}

		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}
	return total

}

func couldBeSubtractive(index int, currenSymbol uint8, roman string) bool {
	isSubstractive := currenSymbol == 'I' || currenSymbol == 'X' || currenSymbol == 'C'
	return index+1 < len(roman) && isSubstractive
}

func (r romanNumerals) ValueOf(symbols ...byte) int {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}
	return 0
}
