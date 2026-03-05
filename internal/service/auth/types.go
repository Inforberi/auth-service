package auth

type RegisterInput struct {
	Email     string
	Password  string
	IP        string
	UserAgent string
}
