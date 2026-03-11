package auth

type registerEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerEmailResponse struct {
	UserID string `json:"user_id"`
}

type LoginEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginEmailResponse struct {
	UserID string `json:"user_id"`
}

type MeResponse struct {
	UserID string `json:"user_id"`
}
