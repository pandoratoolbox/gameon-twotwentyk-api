package handlers

import (
	"gameon-twotwentyk-api/emails"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
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

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	mid := r.Context().Value(models.CTX_user_id).(int64)

	input := struct {
		Code string `json:"code"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	verification, err := store.GetVerification(r.Context(), mid)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if verification.Code != nil {
		if *verification.Code != input.Code && *verification.Method == "Email" {
			ServeError(w, "Incorrect Code", http.StatusBadRequest)
			return
		}
	}

	err = store.DeleteVerification(r.Context(), *verification.Id);

	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	data := models.User{}

	verified := true
	data.VerifiedEmail = &verified

	data.Id = &mid

	err = store.UpdateUser(r.Context(), data)
	
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, map[string]interface{}{
		"success": true,
	})

}

func VerifyEmailByLink(w http.ResponseWriter, r *http.Request) {
	// mid := r.Context().Value(models.CTX_user_id).(int64)
	tokenStr := chi.URLParam(r, "token")

	token, err := jwt.ParseWithClaims(tokenStr, &emails.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("h1l32b"), nil
	})

	if err != nil {
		ServeError(w, "Invalid Token", http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(*emails.CustomClaims)
	if !ok || !token.Valid {
		ServeError(w, "Invalid Token", http.StatusBadRequest)
		return
	}

	userId := claims.UserId
	verficationId := claims.Id

	err = store.DeleteVerification(r.Context(), verficationId);

	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	data := models.User{}

	verified := true
	data.VerifiedEmail = &verified

	data.Id = &userId

	err = store.UpdateUser(r.Context(), data)
	
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, map[string]interface{}{
		"success": true,
	})

}

func VerifyPhoneNumber(w http.ResponseWriter, r *http.Request) {
	mid := r.Context().Value(models.CTX_user_id).(int64)

	input := struct {
		Code string `json:"code"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	verification, err := store.GetVerification(r.Context(), mid)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if verification.Code != nil {
		if *verification.Code != input.Code && *verification.Method == "Phone" {
			ServeError(w, "Incorrect Code", http.StatusBadRequest)
			return
		}
	}

	err = store.DeleteVerification(r.Context(), *verification.Id);

	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	data := models.User{}

	verified := true
	data.VerifiedPhone = &verified

	data.Id = &mid

	err = store.UpdateUser(r.Context(), data)
	
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, map[string]interface{}{
		"success": true,
	})

}