package postgres

import "context"

func (s *AuthStore) CreateUserWithEmailPassword(ctx context.Context, email, emailNorm, passwordHash string) (bool, error)
