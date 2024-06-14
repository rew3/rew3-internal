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
 * Note: it will not touch leading or trailing /.
 */
func CleanDirPath(path string) string {
	if path == "" {
		return path
	}

	// Check for leading and trailing slashes
	hasLeadingSlash := strings.HasPrefix(path, "/")
	hasTrailingSlash := strings.HasSuffix(path, "/")

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
	cleanedPath := strings.Join(cleanedSegments, "/")

	// Add back the leading slash if it was present
	if hasLeadingSlash {
		cleanedPath = "/" + cleanedPath
	}

	// Add back the trailing slash if it was present
	if hasTrailingSlash && cleanedPath != "" {
		cleanedPath = cleanedPath + "/"
	}

	return cleanedPath
}
