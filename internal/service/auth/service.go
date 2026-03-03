package auth

type Service struct {
	repo  Auth
	clock Clock
	hash  PasswordHasher
}

func NewService(repo Auth, clock Clock, hash PasswordHasher) *Service {

	return &Service{clock: clock, hash: hash}
}
