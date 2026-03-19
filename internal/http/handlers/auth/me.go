package auth

import (
	"net/http"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
	"github.com/inforberi/auth-service/internal/http/middleware"
)

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	authCtx, ok := middleware.GetAuthContext(r.Context())
	if !ok {
		helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing auth context")
		return
	}

	helpers.WriteJSON(w, http.StatusOK, MeResponse{UserID: authCtx.UserID})

}
