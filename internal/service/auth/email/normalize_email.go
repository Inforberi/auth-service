package email

import "strings"

func NormalizeEmail(email string) string {
	norm := strings.TrimSpace(strings.ToLower(email))

	return norm
}
