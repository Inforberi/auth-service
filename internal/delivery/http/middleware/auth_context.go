package middleware

import "context"

type AuthContext struct {
	UserID         string
	SessionID      string
	SessionVersion int
	TokenHash      []byte
}

type contextKey string

const authContextKey contextKey = "authContext"

func GetAuthContext(ctx context.Context) (AuthContext, bool) {
	v := ctx.Value(authContextKey)
	if v == nil {
		return AuthContext{}, false
	}

	authCtx, ok := v.(AuthContext)
	return authCtx, ok
}
