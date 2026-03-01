package auth

import "strings"

func Normalize(s string) string {
	norm := strings.TrimSpace(strings.ToLower(s))

	return norm
}
