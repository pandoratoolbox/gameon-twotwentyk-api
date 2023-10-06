package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

const (
	CLAIM_STATUS_PENDING  = int64(0)
	CLAIM_STATUS_APPROVED = int64(1)
	CLAIM_STATUS_REJECTED = int64(2)
)

func ApproveClaim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "claim_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}
	status := CLAIM_STATUS_APPROVED

	c := models.Claim{
		ClaimData: models.ClaimData{
			Id:     &id,
			Status: &status,
		},
	}

	err = store.UpdateClaim(ctx, c)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func RejectClaim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "claim_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}
	status := CLAIM_STATUS_REJECTED

	c := models.Claim{
		ClaimData: models.ClaimData{
			Id:     &id,
			Status: &status,
		},
	}

	err = store.UpdateClaim(ctx, c)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func GetClaim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "claim_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetClaim(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewClaim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	input := struct {
		NftCardPredictionId int64
		NftCardTriggerId    int64
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	spew.Dump(input)

	trigger_card, err := store.GetNftCardTrigger(ctx, input.NftCardTriggerId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prediction_card, err := store.GetNftCardPrediction(ctx, input.NftCardPredictionId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var trigger_in_prediction bool
	if prediction_card.NftCardTriggers == nil {
		ServeError(w, "This prediction card has no triggers", http.StatusInternalServerError)
		return
	}

	for _, v := range *prediction_card.NftCardTriggers {
		if *v.Id == input.NftCardTriggerId {
			trigger_in_prediction = true
		}
	}

	if !trigger_in_prediction {
		ServeError(w, "This trigger card is not in the prediction card", http.StatusInternalServerError)
		return
	}

	if *trigger_card.OwnerId != mid {
		ServeError(w, "You are not the owner of this trigger card", http.StatusInternalServerError)
		return
	}

	if *prediction_card.OwnerId != mid {
		ServeError(w, "You are not the owner of this prediction card", http.StatusInternalServerError)
		return
	}

	if prediction_card.IsClaimed != nil && *prediction_card.IsClaimed {
		ServeError(w, "This prediction card has already been claimed", http.StatusInternalServerError)
		return
	}

	new := models.Claim{
		ClaimData: models.ClaimData{
			ClaimerId:       &mid,
			NftPredictionId: &input.NftCardPredictionId,
			NftTriggerId:    &input.NftCardTriggerId,
		},
	}

	err = store.NewClaim(ctx, &new)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, new)
}

func UpdateClaim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := strconv.ParseInt(chi.URLParam(r, "claim_id"), 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	data := models.Claim{}
	data.Id = &id

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	err = store.UpdateClaim(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteClaim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "claim_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteClaim(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListClaimForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	data, err := store.ListClaimByClaimerId(ctx, mid)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}

func ListClaim(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data, err := store.ListClaim(ctx)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
