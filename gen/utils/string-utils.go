package utils

import "unicode"

/**
 * Capatalize the first letter of given text.
 */
func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s // Return empty string if input is empty
	}
	upperFirst := unicode.ToUpper(rune(s[0]))
	capitalized := string(upperFirst) + s[1:]
	return capitalized
}
