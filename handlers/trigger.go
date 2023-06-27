package handlers

import (
	"gameon-twotwentyk-api/store"
	"net/http"
)

func ListTrigger(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := store.ListTrigger(ctx)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, data)
}
