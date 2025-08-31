package utils

import "strings"

// ConfirmAction asks user for yes/no confirmation
func ConfirmAction(input string) bool {
	return strings.ToLower(strings.TrimSpace(input)) == "yes"
}
