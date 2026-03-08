package session

import "context"

func (s *SessionService) UpdateSessionLastSeen(ctx context.Context, sessionID string) error {
	now := s.clock.Now().UTC()

	if err := s.repo.UpdateSessionLastSeen(ctx, sessionID, now); err != nil {
		return ErrUpdateSessionLastSeen
	}

	return nil
}
