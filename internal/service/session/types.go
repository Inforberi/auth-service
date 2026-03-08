package session

import "time"

type CreateSessionResult struct {
	SessionID string
	Token     string
	ExpiresAt time.Time
}

type GetSessionResult struct {
	SessionID      string
	UserID         string
	SessionVersion int
}
