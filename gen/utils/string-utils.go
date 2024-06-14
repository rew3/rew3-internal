package utils

import (
	"strings"
	"unicode"
)

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

/**
 * Removes redundant slashes from the given path.
 * Note: it will also remove leading and trailing slashes.
 * foo//bar/cat, to foo/bar/cat.
 */
func CleanDirPath(path string) string {
	// Split the path into segments
	segments := strings.Split(path, "/")

	// Filter out empty segments
	var cleanedSegments []string
	for _, segment := range segments {
		if segment != "" {
			cleanedSegments = append(cleanedSegments, segment)
		}
	}

	// Join the cleaned segments with a single slash
	return strings.Join(cleanedSegments, "/")
}
