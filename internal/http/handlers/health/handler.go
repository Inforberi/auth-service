package health

import (
	"net/http"

	"github.com/inforberi/auth-service/internal/http/handlers/helpers"
)

func Health(w http.ResponseWriter, r *http.Request) {
	helpers.WriteJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
