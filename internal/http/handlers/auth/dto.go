package auth

type registerEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerEmailResponse struct {
	UserID    string `json:"user_id"`
	SessionID string `json:"session_id"`
	ExpiresAt string `json:"expires_at"`
}
