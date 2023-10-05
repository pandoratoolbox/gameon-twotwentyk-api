package handlers

import (
	"net/http"
	"strconv"

	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"

	"github.com/go-chi/chi"
	"github.com/pandoratoolbox/json"
)

const PINATA_ACCESS_TOKEN = "DmLxLXKWM6V3doP4C9YfMD0u1mYJvTmsmfHYJfvj9ShHRHVOqWud0w6irfCDGDyo"

func CreateNftCollection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := struct {
		Name string
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	list, err := store.ListCardCollection(ctx)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	name := "Collection " + strconv.Itoa(len(list)+1)

	if input.Name != "" {
		name = input.Name
	}

	card_collection := models.CardCollection{
		CardCollectionData: models.CardCollectionData{
			Name: &name,
		},
	}

	err = store.NewCardCollection(ctx, &card_collection)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	ServeJSON(w, card_collection)
}

func UpdateNftCollectionConfig(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	collection_id, err := strconv.ParseInt(chi.URLParam(r, "card_collection_id"), 10, 64)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	input := struct {
		CardSeries []models.CardSeriesData
	}{}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	var card_series []*models.CardSeries

	for _, v := range input.CardSeries {
		v.CardCollectionId = &collection_id

		o := models.CardSeries{
			CardSeriesData: v,
		}

		err := store.NewCardSeries(ctx, &o)
		if err != nil {
			ServeError(w, err.Error(), 500)
			return
		}

		card_series = append(card_series, &o)

		//send config to engine
		//await engine to produce agg_pack
		//generate nft card packs from agg_pack json
	}

}
