package auth

import (
	"net/mail"
	"unicode"
)

func ValidateEmail(email string) error {
	if email == "" {
		return ErrEmptyEmail
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrPasswordTooShort
	}

	if len(password) > 128 {
		return ErrPasswordTooLong
	}

	var hasLetter bool
	var hasDigit bool

	for _, r := range password {
		if unicode.IsLetter(r) {
			hasLetter = true
		}

		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}

	if !hasLetter {
		return ErrPasswordNoLetter
	}

	if !hasDigit {
		return ErrPasswordNoDigit
	}

	return nil
}
