package function

import (
	"strings"
)

// CheckNotEmpty check if the input string is not empty, return true if not empty
func CheckNotEmpty(s string) bool {
	return strings.Trim(s, " ") != ""
}
