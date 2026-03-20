package email

import (
	"net/http"

	"github.com/inforberi/auth-service/internal/delivery/http/handlers/helpers"
	"github.com/inforberi/auth-service/internal/delivery/http/middleware"
)

func (h *EmailHandler) Me(w http.ResponseWriter, r *http.Request) {
	authCtx, ok := middleware.GetAuthContext(r.Context())
	if !ok {
		helpers.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing auth context")
		return
	}

	helpers.WriteJSON(w, http.StatusOK, MeResponse{UserID: authCtx.UserID})

}
