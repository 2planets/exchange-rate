package utils

import (
	"strconv"
	"strings"
)

// Formats a float to two decimal places
// and adds commas to represent thousands
func FormatFloatWithCommas(val float64) string {
	parseFloat := strconv.FormatFloat(val, 'f', 2, 64)
	dotPos := strings.Index(parseFloat, ".")
	integerPart := parseFloat[:dotPos]
	return AddComma(integerPart) + parseFloat[dotPos:]
}

// Adds commas to represent thousands
// Example: 1,234,567
func AddComma(s string) string {
	var result string
	length := len(s)
	for i := 0; i < length; i++ {
		if i > 0 && (length-i)%3 == 0 {
			result += ","
		}
		result += string(s[i])
	}
	return result
}

// Removes special characters from the source string
func RemoveSpecialChars(source string) string {
	replacer := strings.NewReplacer("$", "", ",", "")
	return replacer.Replace(source)
}

func IsInteger(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}
