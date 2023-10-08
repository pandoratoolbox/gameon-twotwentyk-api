package models

type NftCardDayMonthFilter struct {
	CardCollectionId *int64
	Rarities         *Ints
	Status           *Ints
	Day              *int64
	Month            *int64
}

type NftCardYearFilter struct {
	CardCollectionId *int64
	Year             *int64
	Rarities         *Ints
	Status           *Ints
}

type NftCardCategoryFilter struct {
	CardCollectionId *int64
	Rarities         *Ints
	Status           *Ints
	Categories       *Strings
}

type NftCardTriggerFilter struct {
	Tiers            *Strings
	Rarities         *Ints
	Status           *Ints
	Triggers         *Strings
	CardCollectionId *int64
}

type NftCardIdentityFilter struct {
	CardCollectionId *int64
	Rarities         *Ints
	Status           *Ints
	Categories       *Strings
	Celebrities      *Strings
}

type NftCardPredictionFilter struct {
	Triggers    *Strings
	Celebrities *Strings
	Rarities    *Ints
	Status      *Ints
	// CardSeriesId *int64
	CardCollectionId *int64
}

type NftCardCraftingFilter struct {
	Rarities         *Ints
	Status           *Ints
	CardCollectionId *int64
}
