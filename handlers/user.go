package handlers

import (
	"fmt"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"gameon-twotwentyk-api/twilio"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "user_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetUser(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewUser(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	input := struct {
		PhoneNumber string `json:"phone_number"`
		Username    string `json:"username"`
		Name        string `json:"name"`
		Password    string `json:"password"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	data := models.User{}

	if input.PhoneNumber != "" {
		data.PhoneNumber = &input.PhoneNumber

		user, err := store.GetUser(r.Context(), mid)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		code := fmt.Sprintf("%06d", rand.Intn(999999-100000+1)+100000)
		method := "Phone"

		verification := models.Verification{
			VerificationData: models.VerificationData{
				UserId: user.Id,
				Code:   &code,
				Method: &method,
			},
		}

		err = store.NewVerification(ctx, &verification)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body := fmt.Sprintf("Here's your TwoTwentyK Verification Code: %s", code)
		twilio.SendSMS(*data.PhoneNumber, body)
	}

	if input.Name != "" {
		data.Name = &input.Name
	}

	if input.Username != "" {
		data.Username = &input.Username
	}

	if input.Password != "" {
		data.Password = &input.Password
	}

	data.Id = &mid

	err = store.UpdateUser(ctx, data)
	if err != nil {
		if data.Username != nil {
			ServeError(w, "Username already exists", 400)
			return
		}

		ServeError(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "user_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteUser(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
