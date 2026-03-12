package auth

import (
	"context"
	"time"
)

type RegisterInput struct {
	Email     string
	Password  string
	IP        *string
	UserAgent *string
	DeviceID  *string
}

type RegisterResult struct {
	UserID    string
	SessionID string
	Token     string
	ExpiresAt time.Time
}

func (s *AuthService) RegisterEmail(ctx context.Context, input RegisterInput) (RegisterResult, error) {
	enabled, err := s.repo.IsProviderEnabled(ctx, "email")
	if err != nil {
		return RegisterResult{}, err
	}
	if !enabled {
		return RegisterResult{}, ErrProviderNotEnabled
	}

	email := input.Email
	password := input.Password

	// Normalize Email
	normalizeEmail := NormalizeEmail(email)

	// Validate Email
	if err = ValidateEmail(normalizeEmail); err != nil {
		return RegisterResult{}, err
	}

	// Validate Password
	if err = ValidatePassword(password); err != nil {
		return RegisterResult{}, err
	}

	// Hash password
	hashPassword, err := s.hash.Hash(password)
	if err != nil {
		return RegisterResult{}, err
	}

	now := s.clock.Now().UTC()

	// Create user
	userID, sessionVersion, err := s.repo.CreateUserWithEmailPassword(ctx, email, normalizeEmail, hashPassword, now)
	if err != nil {
		if isRepoEmailTaken(err) {
			return RegisterResult{}, ErrEmailTaken
		}
		return RegisterResult{}, err
	}

	// Create session
	sessionRes, err := s.sessions.CreateSession(
		ctx,
		userID,
		sessionVersion,
		input.IP,
		input.UserAgent,
		input.DeviceID,
	)
	if err != nil {
		return RegisterResult{}, err
	}

	return RegisterResult{
		UserID:    userID,
		SessionID: sessionRes.SessionID,
		Token:     sessionRes.Token,
		ExpiresAt: sessionRes.ExpiresAt,
	}, nil
}
