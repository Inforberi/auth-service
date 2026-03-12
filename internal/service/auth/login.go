package auth

import (
	"context"
	"time"
)

type LoginInput struct {
	Email     string
	Password  string
	IP        *string
	UserAgent *string
	DeviceID  *string
}

type LoginResult struct {
	UserID    string
	SessionID string
	Token     string
	ExpiresAt time.Time
}

func (s *AuthService) LoginWithEmail(ctx context.Context, input LoginInput) (LoginResult, error) {
	email := input.Email
	password := input.Password
	normEmail := NormalizeEmail(email)

	if email == "" || password == "" || len(password) > 1000 {
		return LoginResult{}, ErrInvalidCredentials
	}

	userID, passwordHash, sessionVersion, disabledAt, found, err := s.repo.GetUserByEmail(ctx, normEmail)
	if err != nil {
		return LoginResult{}, ErrLogin
	}

	if !found {
		return LoginResult{}, ErrInvalidCredentials
	}

	if disabledAt != nil {
		return LoginResult{}, ErrUserDisabled
	}

	if !s.hash.Compare(passwordHash, password) {
		return LoginResult{}, ErrInvalidCredentials
	}

	session, err := s.sessions.CreateSession(ctx, userID, sessionVersion, input.IP, input.UserAgent, input.DeviceID)
	if err != nil {
		return LoginResult{}, ErrLogin
	}

	return LoginResult{
		UserID:    userID,
		SessionID: session.SessionID,
		Token:     session.Token,
		ExpiresAt: session.ExpiresAt,
	}, nil

}
