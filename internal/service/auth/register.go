package auth

import "context"

func (s *Service) RegisterEmail(ctx context.Context, input RegisterInput) error {
	enabled, err := s.repo.IsProviderEnabled(ctx, "email")
	if err != nil {
		return err
	}
	if !enabled {
		return ErrProviderNotEnabled
	}

	email := input.Email
	password := input.Password

	// Normalize Email
	normalizeEmail := Normalize(email)

	// Validate Email
	if err = ValidateEmail(normalizeEmail); err != nil {
		return err
	}

	// Validate Password
	if err = ValidatePassword(password); err != nil {
		return err
	}

	// Hash password
	hashPassword, err := s.hash.Hash(password)
	if err != nil {
		return err
	}

	// Create user
	userId, err := s.repo.CreateUserWithEmailPassword(ctx, email, normalizeEmail, hashPassword, s.clock.Now())
	if err != nil {
		return err
	}

	// TODO: Create session 
	s.repo.CreateSession(ctx, userId)
}
