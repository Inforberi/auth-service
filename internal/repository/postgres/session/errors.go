package session

type sessionNotFoundError struct{}

func (sessionNotFoundError) Error() string {
	return "session not found"
}

func (sessionNotFoundError) SessionNotFound() bool {
	return true
}

var ErrSessionNotFound error = sessionNotFoundError{}

type userNotFoundError struct{}

func (userNotFoundError) Error() string {
	return "user not found"
}

func (userNotFoundError) UserNotFound() bool {
	return true
}

var ErrUserNotFound error = userNotFoundError{}
