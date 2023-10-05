package models

import "time"

// type CollectionCardPackConfig map[string]map[int64]CardPackConfig //tier - card pack id - card pack config
type CardSeriesCardPackConfig map[int64]CardPackConfig

// used to generate the collection - saved in database and factory contract

type CardCollectionData struct {
	Id                           *int64
	DayMonthAddress              *string
	YearAddress                  *string
	AggPackPath                  *string
	TriggerAddress               *string
	CardPackAddress              *string
	CraftingAddress              *string
	Name                         *string
	IdentityAddress              *string
	PredictionAddress            *string
	Status                       *int64
	CategoryAddress              *string
	TriggerPrizePoolDistribution *TriggerPrizePoolDistribution // trigger tier - amount
	Rarities                     *Strings
	Tiers                        *Strings
	// CollectionCards              CollectionCardPackConfig
	Images    CardCollectionImages
	CreatedAt *time.Time
}

type CardCollectionImages struct {
	Logo       *string
	TierImages *map[int64]struct {
		CardBack *string
	} //tier - image
	CardImages *map[int64]map[int64]struct {
		CardFront *string
		Animation *string
	} //card type - rarity - image
}

type TriggerPrizePoolDistribution map[string]int64

type CardSeriesData struct {
	CardsPerPack     *int64
	CardPackQuantity *int64
	CostUsd          *int64
	Guaranteed       *map[int64]CardTypeValue //rarity - card type - quantity
	Probabilities    struct {
		Year     RarityValues
		DayMonth RarityValues
		Category RarityValues
		Trigger  map[string]map[int64]struct {
			Probability   float64
			RewardType    int64
			RewardRarity  int64
			RegDefinition float64
		} //trigger tier - rarity - trigger reward config
	}
	Id               *int64
	CardCollectionId *int64
	Name             *string
	CardPacks        CardSeriesCardPackConfig
	// Images           CardSeriesMetadataImages
}

type CardSeriesMetadataImages struct {
	CardBack map[int64]string
} //rarity - image

type CardImages struct {
	CardFront string
	CardBack  string
	Animation string
}

type RarityValues map[int64]float64

type CardTypeValue struct {
	Category int64
	DayMonth int64
	Trigger  []struct {
		Name string
		Tier string
		Id   int64
	}
	Year int64
}

// config for each card pack
type CardPackConfig struct {
	Changed    int
	CardPackId int64
	Tier       string
	Contains   struct {
		Rare     CardTypeValue
		Core     CardTypeValue
		Uncommon CardTypeValue
	}
}
