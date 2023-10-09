package handlers

import (
	"fmt"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

func GetNftCardIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_identity_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	data, err := store.GetNftCardIdentity(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, data)
}

func NewNftCardIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := models.NftCardIdentity{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = store.NewNftCardIdentity(ctx, &input)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ServeJSON(w, input)
}

func UpdateNftCardIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_identity_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	input := struct {
		CelebrityId int64
	}{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	celebrity, err := store.GetCelebrity(ctx, input.CelebrityId)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	identity_card, err := store.GetNftCardIdentity(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	spew.Dump(identity_card)

	spew.Dump(celebrity)

	if *identity_card.CardSeries.CardCollectionId != *celebrity.CardCollectionId {
		ServeError(w, "Card collection doesn't match", 400)
		return
	}

	if *identity_card.Day != *celebrity.BirthDay {
		ServeError(w, "Day doesn't match", 400)
		return
	}

	if *identity_card.Month != *celebrity.BirthMonth {
		ServeError(w, "Month doesn't match", 400)
		return
	}

	if *identity_card.Year != *celebrity.BirthYear {
		ServeError(w, "Year doesn't match", 400)
		return
	}

	if *identity_card.Category != *celebrity.Category {
		ServeError(w, "Category doesn't match", 400)
		return
	}

	data := models.NftCardIdentity{
		NftCardIdentityData: models.NftCardIdentityData{
			Id:            &id,
			CelebrityName: celebrity.Name,
		},
	}

	err = store.UpdateNftCardIdentity(ctx, data)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func DeleteNftCardIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q := chi.URLParam(r, "nft_card_identity_id")
	id, err := strconv.ParseInt(q, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	err = store.DeleteNftCardIdentity(ctx, id)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}

func ListNftCardIdentityForUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mid := ctx.Value(models.CTX_user_id).(int64)

	filters := models.NftCardIdentityFilter{}

	q_card_collection_id := r.URL.Query().Get("card_collection_id")
	if q_card_collection_id != "" {
		i, err := strconv.ParseInt(q_card_collection_id, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}

		filters.CardCollectionId = &i
	}

	q_rarities := r.URL.Query().Get("rarities")
	if q_rarities != "" {
		var rarities models.Ints
		for _, v := range strings.Split(q_rarities, ",") {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), 500)
				return
			}

			rarities = append(rarities, i)
		}

		filters.Rarities = &rarities
	}

	var categories models.Strings
	categories_raw := r.URL.Query().Get("categories")
	if categories_raw != "" {
		categories_str := strings.Split(categories_raw, ",")
		for _, v := range categories_str {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			c, ok := store.CategoryMap[i]
			if !ok {
				ServeError(w, fmt.Sprintf("Invalid category id: %d", i), http.StatusBadRequest)
				return
			}
			categories = append(categories, *c.Name)
		}
		filters.Categories = &categories
	}

	var celebrities models.Strings
	celebrities_raw := r.URL.Query().Get("celebrities")
	if celebrities_raw != "" {
		celebrities_str := strings.Split(celebrities_raw, ",")
		for _, v := range celebrities_str {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			c, ok := store.CelebrityMap[i]
			if !ok {
				ServeError(w, fmt.Sprintf("Invalid category id: %d", i), http.StatusBadRequest)
				return
			}
			celebrities = append(celebrities, *c.Name)
		}
		filters.Celebrities = &celebrities
	}

	q_status := r.URL.Query().Get("status")
	if q_status != "" {
		var status models.Ints
		for _, v := range strings.Split(q_status, ",") {
			i, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				ServeError(w, err.Error(), 500)
				return
			}

			status = append(status, i)
		}

		filters.Status = &status
	}

	data, err := store.ListNftCardIdentityByOwnerId(ctx, mid, filters)
	if err != nil {
		ServeError(w, err.Error(), 400)
		return
	}

	ServeJSON(w, data)
}
