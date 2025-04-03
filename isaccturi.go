package accturi

import (
	"strings"
	"unsafe"
)

// IsAcctURIBytes returns true if `value` is a acct-uri, and returns false otherwise.
func IsAcctURIBytes(value []byte) bool {
	var str string = unsafe.String(unsafe.SliceData(value), len(value))

	return IsAcctURIString(str)
}


// IsAcctURIString returns true if `value` is a acct-uri, and returns false otherwise.
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
