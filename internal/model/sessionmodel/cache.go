package sessionmodel

type CacheSession struct {
	SessionID      string `json:"session_id"`
	UserID         string `json:"user_id"`
	SessionVersion int    `json:"session_version"`
	ExpiresAtUnix  int64  `json:"expires_at_unix"`
	Revoked        bool   `json:"revoked"`
	UserDisabled   bool   `json:"user_disabled"`
}
