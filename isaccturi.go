package accturi

import (
	"strings"
)

// IsAcctURIString returns true is `value` is a acct-uri, and returns false otherwise.
func IsAcctURIString(value string) bool {

	if !strings.HasPrefix(value, prefix) {
		return false
	}

	var length int = len(value)

	if length <= len(prefix) {
		return false
	}

	return true
}
