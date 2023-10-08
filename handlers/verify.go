package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"

	"github.com/pandoratoolbox/json"
)

func VerifyPassword(w http.ResponseWriter, r *http.Request) {
	mid := r.Context().Value(models.CTX_user_id).(int64)

	input := struct {
		Password string `json:"password"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := store.GetUser(r.Context(), mid)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Password != nil {
		if *user.Password != input.Password {
			ServeError(w, "Incorrect password", http.StatusBadRequest)
			return
		}
	}

	ServeJSON(w, map[string]interface{}{
		"success": true,
	})

}
