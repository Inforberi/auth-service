package sessionredis

type SessionSnapshot struct {
	SessionID      string `json:"session_id"`
	UserID         string `json:"user_id"`
	SessionVersion int    `json:"session_version"`
	ExpiresAtUnix  int64  `json:"expires_at_unix"`
	Revoked        bool   `json:"revoked"`
}

type userSessionVersion struct {
	ActualSessionVersion int `json:"actual_session_version"`
}
