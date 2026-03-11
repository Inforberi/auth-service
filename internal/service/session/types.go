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

type GetRepoResult struct {
	SessionID            string
	UserID               string
	SessionVersion       int
	ExpiresAt            time.Time
	RevokedAt            *time.Time
	ActualSessionVersion int
	DisabledAt           *time.Time
	Found                bool
}
