package auth

import "time"

// register
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

// Login
type LoginInput struct {
	Email     string
	Password  string
	IP        *string
	UserAgent *string
	DeviceID  *string
}

type LoginResult struct {
	UserID    string
	SessionID string
	Token     string
	ExpiresAt time.Time
}
