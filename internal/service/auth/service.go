package auth

type Service struct {
	clock Clock
	hash  PasswordHasher
}

func NewService(clock Clock, hash PasswordHasher) *Service {

	return &Service{clock: clock, hash: hash}
}
