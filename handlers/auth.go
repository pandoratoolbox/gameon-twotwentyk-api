package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"

	"github.com/Timothylock/go-signin-with-apple/apple"
	"github.com/pandoratoolbox/json"

	"google.golang.org/api/idtoken"
)

const (
	GOOGLE_CLIENT_ID = "620329827727-t3sttbu6556u69ebv50fmt5rda85drp0.apps.googleusercontent.com"
)

func AuthApple(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := struct {
		IdToken string
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.IdToken == "" {
		ServeError(w, "Missing id_token", http.StatusBadRequest)
		return
	}

	claims, err := apple.GetClaims(input.IdToken)
	if err != nil {
		ServeError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	r_email, ok := claims.Get("email")
	if !ok {
		ServeError(w, "Error getting email from google id token claims", http.StatusInternalServerError)
		return
	}

	email := r_email.(string)

	user, err := store.GetUserByEmail(ctx, email)
	if err != nil {
		// ServeError(w, err.Error(), http.StatusInternalServerError)
		// return
		roles := models.Ints{1}

		user := models.User{
			UserData: models.UserData{
				Email:   &email,
				RoleIds: &roles,
			},
		}

		err = store.NewUser(ctx, &user)
		if err != nil {
			ServeError(w, err.Error(), 400)
			return
		}

		store.AddNftsToUser(ctx, *user.Id)

		_, jwtstring, err := TokenAuth.Encode(map[string]interface{}{
			"id":       *user.Id,
			"role_ids": *user.RoleIds,
		})

		if err != nil {
			ServeError(w, err.Error(), 400)
			return
		}

		response := struct {
			User  models.User
			Token string
		}{
			User:  user,
			Token: jwtstring,
		}

		ServeJSON(w, response)
		return
	}

	_, jwtstring, err := TokenAuth.Encode(map[string]interface{}{
		"id":       *user.Id,
		"role_ids": *user.RoleIds,
	})
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	out := struct {
		Token string
	}{
		Token: jwtstring,
	}

	ServeJSON(w, out)
}

func AuthGoogle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := struct {
		IdToken string
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.IdToken == "" {
		ServeError(w, "Missing id_token", http.StatusBadRequest)
		return
	}

	res, err := idtoken.Validate(ctx, input.IdToken, GOOGLE_CLIENT_ID)
	if err != nil {
		ServeError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	r_email, ok := res.Claims["email"]
	if !ok {
		ServeError(w, "Error getting email from google id token claims", http.StatusInternalServerError)
		return
	}

	email := r_email.(string)

	user, err := store.GetUserByEmail(ctx, email)
	if err != nil {
		// ServeError(w, err.Error(), http.StatusInternalServerError)
		// return
		roles := models.Ints{1}

		user := models.User{
			UserData: models.UserData{
				Email:   &email,
				RoleIds: &roles,
			},
		}

		err = store.NewUser(ctx, &user)
		if err != nil {
			ServeError(w, err.Error(), 400)
			return
		}

		store.AddNftsToUser(ctx, *user.Id)

		_, jwtstring, err := TokenAuth.Encode(map[string]interface{}{
			"id":       *user.Id,
			"role_ids": *user.RoleIds,
		})

		if err != nil {
			ServeError(w, err.Error(), 400)
			return
		}

		response := struct {
			User  models.User
			Token string
		}{
			User:  user,
			Token: jwtstring,
		}

		ServeJSON(w, response)
		return
	}

	_, jwtstring, err := TokenAuth.Encode(map[string]interface{}{
		"id":       *user.Id,
		"role_ids": *user.RoleIds,
	})
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	out := struct {
		Token string
	}{
		Token: jwtstring,
	}

	ServeJSON(w, out)
}
