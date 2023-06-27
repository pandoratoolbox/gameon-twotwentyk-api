package handlers

import (
	"context"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"sort"
	"time"
)

func CraftIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	mid := ctx.Value(models.CTX_user_id).(int64)

	input := struct {
		NftCardCraftingId int64 `json:"nft_card_crafting_id"`
		NftCardDayMonthId int64 `json:"nft_card_day_month_id"`
		NftCardYearId     int64 `json:"nft_card_year_id"`
		NftCardCategoryId int64 `json:"nft_card_category_id"`
		CelebrityId       int64 `json:"celebrity_id"`
	}{}

	// var out models.NftCardIdentity

	nft_card_day_month, err := store.GetNftCardDayMonth(ctx, input.NftCardDayMonthId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_year, err := store.GetNftCardYear(ctx, input.NftCardYearId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_crafting, err := store.GetNftCardCrafting(ctx, input.NftCardCraftingId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_category, err := store.GetNftCardCategory(ctx, input.NftCardCategoryId)
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

	date_from_nfts := time.Date(int(*nft_card_year.Year), time.Month(*nft_card_day_month.Month), int(*nft_card_day_month.Day), 0, 0, 0, 0, nil)

	celebrity_birthdate := time.Date(int(*celebrity.BirthYear), time.Month(*celebrity.BirthMonth), int(*celebrity.BirthDay), 0, 0, 0, 0, nil)

	if date_from_nfts != celebrity_birthdate {
		ServeError(w, "NFT date combination and celebrity birthdate doesn't match", 400)
		return
	}

	day := nft_card_day_month.Day
	month := nft_card_day_month.Month
	year := nft_card_year.Year
	rarities := make(map[int64]int)

	_, ok := rarities[*nft_card_day_month.Rarity]
	if !ok {
		rarities[*nft_card_day_month.Rarity] = 1
	} else {
		rarities[*nft_card_day_month.Rarity]++
	}

	_, ok = rarities[*nft_card_year.Rarity]
	if !ok {
		rarities[*nft_card_year.Rarity] = 1
	} else {
		rarities[*nft_card_year.Rarity]++
	}

	_, ok = rarities[*nft_card_category.Rarity]
	if !ok {
		rarities[*nft_card_category.Rarity] = 1
	} else {
		rarities[*nft_card_category.Rarity]++
	}

	rarity := int64(0)
	var rar_arr []int
	for k, v := range rarities {
		rar_arr = append(rar_arr, int(k))
		if k == *nft_card_year.Rarity {
			if v >= 2 {
				rarity = k
			}
		}
	}

	sort.Ints(rar_arr)
	lowest_rarity := rar_arr[0]

	if int(rarity)-lowest_rarity > 1 {
		rarity = 0
	}

	c_name := celebrity.Name
	category := nft_card_category.Category
	is_crafted := false

	out := models.NftCardIdentity{
		NftCardIdentityData: models.NftCardIdentityData{
			OwnerId:       &mid,
			Year:          year,
			Month:         month,
			Day:           day,
			Rarity:        &rarity,
			CelebrityName: c_name,
			Category:      category,
			IsCrafted:     &is_crafted,
		},
	}

	err = store.NewNftCardIdentity(context.TODO(), &out)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	burn := true

	nft_card_category.IsCrafted = &burn
	store.UpdateNftCardCategory(ctx, nft_card_category)

	nft_card_crafting.IsCrafted = &burn
	store.UpdateNftCardCrafting(ctx, nft_card_crafting)

	nft_card_day_month.IsCrafted = &burn
	store.UpdateNftCardDayMonth(ctx, nft_card_day_month)

	nft_card_year.IsCrafted = &burn
	store.UpdateNftCardYear(ctx, nft_card_year)

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

	mid := ctx.Value(models.CTX_user_id).(int64)

	var out models.NftCardPrediction

	input := struct {
		NftCardCraftingId int64   `json:"nft_card_crafting_id"`
		NftCardIdentityId int64   `json:"nft_card_identity_id"`
		NftCardTriggerIds []int64 `json:"nft_card_trigger_ids"`
	}{}

	nft_card_crafting, err := store.GetNftCardCrafting(ctx, input.NftCardCraftingId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nft_card_identity, err := store.GetNftCardIdentity(ctx, input.NftCardIdentityId)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var triggers []string
	var nft_card_triggers []models.NftCardTrigger
	trigger_highest_rarity := int64(0)
	for _, t := range input.NftCardTriggerIds {
		nft_card_trigger, err := store.GetNftCardTrigger(ctx, t)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nft_card_triggers = append(nft_card_triggers, nft_card_trigger)

		if *nft_card_trigger.OwnerId != mid {
			ServeError(w, "User does not own trigger card", 400)
			return
		}

		if *nft_card_trigger.Rarity > trigger_highest_rarity {
			trigger_highest_rarity = *nft_card_trigger.Rarity
		}

		triggers = append(triggers, *nft_card_trigger.Trigger)
	}

	rarity := trigger_highest_rarity

	is_claimed := false

	ttriggers := models.Strings(triggers)

	out = models.NftCardPrediction{
		NftCardPredictionData: models.NftCardPredictionData{
			OwnerId:       &mid,
			IsClaimed:     &is_claimed,
			Rarity:        &rarity,
			Triggers:      &ttriggers,
			CelebrityName: nft_card_identity.CelebrityName,
		},
	}

	err = store.NewNftCardPrediction(ctx, &out)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	burn := true

	nft_card_identity.IsCrafted = &burn
	store.UpdateNftCardIdentity(ctx, nft_card_identity)

	nft_card_crafting.IsCrafted = &burn
	store.UpdateNftCardCrafting(ctx, nft_card_crafting)

	for _, t := range nft_card_triggers {
		t.IsCrafted = &burn
		store.UpdateNftCardTrigger(ctx, t)
	}

	ServeJSON(w, out)
}
