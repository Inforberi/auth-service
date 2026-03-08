package auth

import "time"

type RegisterInput struct {
	Email     string
	Password  string
	IP        *string
	UserAgent *string
	DeviceID  *string
}

type RegisterResult struct {
	UserID    string
	SessionID string
	Token     string
	ExpiresAt time.Time
}
