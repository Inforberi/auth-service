package postgres

type emailTakenError struct{}

func (emailTakenError) Error() string {
	return "email already taken"
}

func (emailTakenError) EmailTaken() bool {
	return true
}

var ErrEmailTaken error = emailTakenError{}
