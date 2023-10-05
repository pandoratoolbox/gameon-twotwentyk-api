package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"gameon-twotwentyk-api/handlers"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

type CollectionTierProbabilities struct {
	DayMonth RarityValues
	Year     RarityValues
	Trigger  RarityValues
	Category RarityValues
}

type CollectionTierGuaranteed struct {
	DayMonth RarityValues
	Year     RarityValues
	Trigger  RarityValues
	Category RarityValues
}

type TierConfig struct {
	Price         string
	PackQuantity  string
	Name          string
	Probabilities CollectionTierProbabilities
	Guaranteed    CollectionTierGuaranteed
	CardsPerPack  int
}

type TriggerListRow struct {
	Name string
	Tier string
}

type TriggerList []TriggerListRow

type CollectionMetadata struct {
	CategoryAddress   string
	TriggerAddress    string
	DayMonthAddress   string
	YearAddress       string
	CardPackAddress   string
	CraftingAddress   string
	IdentityAddress   string
	PredictionAddress string
	Name              string
	Id                int64
}

type CollectionConfig struct {
	Metadata     CollectionMetadata
	Tiers        []TierConfig
	AggPackPath  string
	TriggerList  TriggerList
	IdentityList []struct {
		Name      string
		Category  string
		Birthdate string
	}
	TriggerPrizePool struct {
		Minor1 int
		Minor2 int
		Major  int
	}
}

type RarityValues struct {
	Rare     int
	Uncommon int
	Core     int
}

func ClearCollectionMetadata(w http.ResponseWriter, r *http.Request) {}

func UploadCollectionMetadata(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collection_id, err := strconv.ParseInt(chi.URLParam(r, "card_collection_id"), 10, 64)
	if err != nil {
		handlers.ServeError(w, "url id error: "+err.Error(), 400)
		return
	}

	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		handlers.ServeError(w, err.Error(), 400)
		return
	}

	triggers_file, _, err := r.FormFile("triggers")
	if err != nil {
		handlers.ServeError(w, err.Error(), 400)
		return
	}

	defer triggers_file.Close()

	reader := csv.NewReader(triggers_file)

	fmt.Println("started reading trigger csv")

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		trigger_tier := line[0]
		trigger_name := line[1]

		new := models.Trigger{
			TriggerData: models.TriggerData{
				Name:             &trigger_name,
				Tier:             &trigger_tier,
				CardCollectionId: &collection_id,
			},
		}

		err = store.NewTrigger(context.Background(), &new)
		if err != nil {
			log.Println(err)
		}
	}

	celebrities_file, _, err := r.FormFile("celebrities")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer triggers_file.Close()

	c_reader := csv.NewReader(celebrities_file)

	for {
		line, err := c_reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			handlers.ServeError(w, err.Error(), 500)
			return
		}

		name := line[0]
		category := line[1]
		birthdate := line[2]
		spl_birthdate := strings.Split(birthdate, "-")
		if len(spl_birthdate) != 3 {
			http.Error(w, "Invalid birthdate format", http.StatusInternalServerError)
			return
		}

		birth_day, err := strconv.ParseInt(spl_birthdate[1], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		birth_month, err := strconv.ParseInt(spl_birthdate[0], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		birth_year, err := strconv.ParseInt(spl_birthdate[2], 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		new := models.Celebrity{
			CelebrityData: models.CelebrityData{
				Name:             &name,
				Category:         &category,
				BirthDay:         &birth_day,
				BirthMonth:       &birth_month,
				BirthYear:        &birth_year,
				CardCollectionId: &collection_id,
			},
		}

		err = store.NewCelebrity(context.Background(), &new)
		if err != nil {
			log.Println(err)
		}

	}

	status := int64(1)

	err = store.UpdateCardCollection(ctx, models.CardCollection{
		CardCollectionData: models.CardCollectionData{
			Id:     &collection_id,
			Status: &status,
		}})

	if err != nil {
		handlers.ServeError(w, err.Error(), 400)
		return
	}

	w.WriteHeader(200)
}

func ImportMetadata() error {
	//read csv for trigger list
	//read csv for identities
	//get collection metadata from post body
	// triggers, err := ImportTriggers(trigger_path)
	// if err != nil {
	// 	return err
	// }

	// celebrities, err := ImportCelebrities(celebrity_path)
	// if err != nil {
	// 	return err
	// }

	return nil
}
