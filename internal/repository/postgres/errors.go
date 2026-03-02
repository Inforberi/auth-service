package postgres

import "errors"

var ErrEmailTaken = errors.New("email already taken")
