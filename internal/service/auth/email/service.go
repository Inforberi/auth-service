package email

type EmailService struct {
	repo     Email
	clock    Clock
	hash     PasswordHasher
	sessions SessionCreator
}

func New(repo Email, clock Clock, hash PasswordHasher, sessions SessionCreator) *EmailService {

	return &EmailService{repo: repo, clock: clock, hash: hash, sessions: sessions}
}
