package handlers

import (
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/nft"
	"gameon-twotwentyk-api/store"
	"net/http"
	"time"
)

func CraftIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := struct {
		NftCardDayId   int64
		NftCardMonthId int64
		NftCardYearId  int64
		CelebrityId    int64
	}{}

	var out models.NftCardIdentity

	//get nft card by id
	nft_card_day, err := nft.GetNftCardDay(ctx, input.NftCardDayId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_month, err := nft.GetNftCardMonth(ctx, input.NftCardMonthId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_year, err := nft.GetNftCardYear(ctx, input.NftCardYearId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//get celebrity by id
	celebrity, err := store.GetCelebrity(ctx, input.CelebrityId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	date_from_nfts := time.Date(nft_card_year.Year, time.Month(nft_card_month.Month), nft_card_day.Day, 0, 0, 0, 0, nil)

	celebrity_day := celebrity.Birthdate.Day()
	celebrity_month := celebrity.Birthdate.Month()
	celebrity_year := celebrity.Birthdate.Year()

	celebrity_birthdate := time.Date(celebrity_year, celebrity_month, celebrity_day, 0, 0, 0, 0, nil)

	if date_from_nfts != celebrity_birthdate {
		ServeError(w, "NFT date combination and celebrity birthdate doesn't match", 400)
		return
	}

	ServeJSON(w, out)
}

// func CraftDate(w http.ResponseWriter, r *http.Request) {
// 	var out models.NftCardDate

// 	input := struct {
// 		NftCardDayId   int64
// 		NftCardMonthId int64
// 		NftCardYearId  int64
// 	}{}

// 	ServeJSON(out)
// }

func CraftPrediction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var out models.NftCardPrediction

	input := struct {
		NftCardIdentityId int64
		NftCardTriggerId  int64
	}{}

	nft_card_identity, err := nft.GetNftCardIdentity(ctx, input.NftCardIdentityId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_trigger, err := nft.GetNftCardTrigger(ctx, input.NftCardTriggerId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out = models.NftCardPrediction{
		EventName:     nft_card_trigger.EventName,
		CelebrityName: nft_card_identity.CelebrityName,
	}

	ServeJSON(w, out)
}
