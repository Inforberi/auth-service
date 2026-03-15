package session

import "context"

func (s *SessionService) UpdateSessionActivity(ctx context.Context, sessionID string) error {
	now := s.clock.Now().UTC()
	expiresAt := now.Add(s.sessionTTL)
	threshold := now.Add(-s.activityUpdateInterval)

	if err := s.repo.UpdateSessionActivity(ctx, sessionID, now, expiresAt, threshold); err != nil {
		return ErrUpdateSessionLastSeen
	}

	return nil
}
