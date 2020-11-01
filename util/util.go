package util

import "unicode/utf8"

// Ellipsis will make strings longer than some limit L end in an ellipsis
func Ellipsis(s string, limit int) string {
	if utf8.RuneCountInString(s) <= limit {
		return s
	}

	return string([]rune(s)[:limit]) + "…"

}
