package auth

type AuthService struct {
	repo     Auth
	clock    Clock
	hash     PasswordHasher
	sessions SessionCreator
}

func NewAuthService(repo Auth, clock Clock, hash PasswordHasher, sessions SessionCreator) *AuthService {

	return &AuthService{repo: repo, clock: clock, hash: hash, sessions: sessions}
}
