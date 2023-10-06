package models

import (
	"time"
)

type Strings []string
type Ints []int64
type ctxkey int64

const (
	CTX_is_auth       = ctxkey(0)
	CTX_user_id       = ctxkey(1)
	CTX_user_role_ids = ctxkey(2)
	CTX_user_timezone = ctxkey(3)
)

type ClaimData struct {
	Status          *int64
	CreatedAt       *time.Time
	ClaimerId       *int64
	NftPredictionId *int64
	NftTriggerId    *int64
	ArticleId       *int64
	Id              *int64
}

type Claim struct {
	ClaimData
	Claimer       *User
	Article       *Article
	NftPrediction *NftCardPrediction
	NftTrigger    *NftCardTrigger
}

type NftCardTriggerData struct {
	Rarity       *int64
	Tier         *string
	Id           *int64
	IsCrafted    *bool
	OwnerId      *int64
	Trigger      *string
	CardSeriesId *int64
	CreatedAt    *time.Time
}

type NftCardTrigger struct {
	NftCardTriggerData
	Owner              *User
	MarketplaceListing []*MarketplaceListing
	CardSeries         *CardSeries
}

type CardSeries struct {
	CardSeriesData
	CardPack          []*CardPack
	NftCardCrafting   []*NftCardCrafting
	NftCardCategory   []*NftCardCategory
	NftCardDayMonth   []*NftCardDayMonth
	NftCardPrediction []*NftCardPrediction
	NftCardIdentity   []*NftCardIdentity
	NftCardTrigger    []*NftCardTrigger
	NftCardYear       []*NftCardYear
	CardCollection    *CardCollection
}

type NftCardDayMonthData struct {
	Day          *int64
	Month        *int64
	OwnerId      *int64
	Rarity       *int64
	CardSeriesId *int64
	Id           *int64
	IsCrafted    *bool
	CreatedAt    *time.Time
}

type NftCardDayMonth struct {
	NftCardDayMonthData
	Owner              *User
	MarketplaceListing []*MarketplaceListing
	CardSeries         *CardSeries
}

type NftCardYearData struct {
	OwnerId      *int64
	IsCrafted    *bool
	Id           *int64
	Year         *int64
	Rarity       *int64
	CardSeriesId *int64
	CreatedAt    *time.Time
}

type NftCardYear struct {
	NftCardYearData
	MarketplaceListing []*MarketplaceListing
	Owner              *User
	CardSeries         *CardSeries
}

type UserData struct {
	Username       *string
	CreatedAt      *time.Time
	RoleIds        *Ints
	Name           *string
	Balance        *int64
	Id             *int64
	Password       *string
	ExternalAuthId *string
	WalletAddress  *string
	VenlyId        *string
	Email          *string
	PhoneNumber    *string
}

type User struct {
	UserData
	NftCardYear        []*NftCardYear
	NftCardCrafting    []*NftCardCrafting
	CardPack           []*CardPack
	Claim              []*Claim
	NftCardTrigger     []*NftCardTrigger
	NftCardPrediction  []*NftCardPrediction
	NftCardIdentity    []*NftCardIdentity
	MarketplaceOffer   []*MarketplaceOffer
	MarketplaceListing []*MarketplaceListing
	NftCardDayMonth    []*NftCardDayMonth
	NftCardCategory    []*NftCardCategory
}

type CardCollection struct {
	CardCollectionData
	Trigger    []*Trigger
	Celebrity  []*Celebrity
	CardSeries []*CardSeries
}

type MarketplaceOfferData struct {
	Amount               *int64
	BuyerId              *int64
	Id                   *int64
	Status               *int64
	MarketplaceListingId *int64
}

type MarketplaceOffer struct {
	MarketplaceOfferData
	MarketplaceListing *MarketplaceListing
	Buyer              *User
}

type NftCardIdentityData struct {
	IsCrafted     *bool
	OwnerId       *int64
	CelebrityName *string
	Year          *int64
	Category      *string
	Rarity        *int64
	CardSeriesId  *int64
	Id            *int64
	Month         *int64
	Day           *int64
	CreatedAt     *time.Time
}

type NftCardIdentity struct {
	NftCardIdentityData
	MarketplaceListing []*MarketplaceListing
	CardSeries         *CardSeries
	Owner              *User
}

type TriggerData struct {
	CardCollectionId *int64
	Id               *int64
	Name             *string
	Tier             *string
	// RegDefinition    *float64
	// RewardType       *int64
}

type Trigger struct {
	TriggerData
	CardCollection *CardCollection
}

type NftCardCraftingData struct {
	CardSeriesId *int64
	Id           *int64
	IsCrafted    *bool
	OwnerId      *int64
	Rarity       *int64
	CreatedAt    *time.Time
}

type NftCardCrafting struct {
	NftCardCraftingData
	Owner              *User
	MarketplaceListing []*MarketplaceListing
	CardSeries         *CardSeries
}

type NftCardPredictionData struct {
	IsClaimed       *bool
	NftCardTriggers *[]NftCardTriggerData
	Triggers        *Strings
	Id              *int64
	CelebrityName   *string
	OwnerId         *int64
	Rarity          *int64
	CardSeriesId    *int64
	CreatedAt       *time.Time
}

type NftCardPrediction struct {
	NftCardPredictionData
	Owner              *User
	MarketplaceListing []*MarketplaceListing
	CardSeries         *CardSeries
}

type ArticleSourceData struct {
	Id   *int64
	Name *string
}

type ArticleSource struct {
	ArticleSourceData
	Article []*Article
}

type CelebrityData struct {
	Id               *int64
	Name             *string
	BirthDay         *int64
	BirthMonth       *int64
	BirthYear        *int64
	Category         *string
	EligibleTriggers *Strings
	CardCollectionId *int64
}

type Celebrity struct {
	CelebrityData
	CardCollection *CardCollection
}

type CategoryData struct {
	Id   *int64
	Name *string
}

type Category struct {
	CategoryData
}

type CardPackData struct {
	Id           *int64
	OwnerId      *int64
	IsOpened     *bool
	CardSeriesId *int64
	Cards        *CardPackCards
	Tier         *int64
	CreatedAt    *time.Time
}

type CardPack struct {
	CardPackData
	CardSeries         *CardSeries
	Owner              *User
	MarketplaceListing []*MarketplaceListing
}

type NftCardCategoryData struct {
	Category     *string
	OwnerId      *int64
	Id           *int64
	IsCrafted    *bool
	Rarity       *int64
	CardSeriesId *int64
	CreatedAt    *time.Time
}

type NftCardCategory struct {
	NftCardCategoryData
	CardSeries         *CardSeries
	Owner              *User
	MarketplaceListing []*MarketplaceListing
}

type MarketplaceListingData struct {
	Id                  *int64
	OwnerId             *int64
	NftCardCraftingId   *int64
	NftCardCategoryId   *int64
	NftCardTriggerId    *int64
	CardPackId          *int64
	Price               *int64
	IsListed            *bool
	NftCardDayMonthId   *int64
	NftCardYearId       *int64
	NftCardIdentityId   *int64
	CreatedAt           *time.Time
	NftCardPredictionId *int64
	NftTypeId           *int64
}

type MarketplaceListing struct {
	MarketplaceListingData
	NftCardCrafting   *NftCardCrafting
	NftCardPrediction *NftCardPrediction
	Owner             *User
	NftCardIdentity   *NftCardIdentity
	NftCardYear       *NftCardYear
	NftCardCategory   *NftCardCategory
	NftCardDayMonth   *NftCardDayMonth
	NftCardTrigger    *NftCardTrigger
	MarketplaceOffer  []*MarketplaceOffer
	CardPack          *CardPack
}

type IdentityData struct {
	BirthMonth *int64
	BirthYear  *int64
	Category   *string
	Id         *int64
	Name       *string
	BirthDay   *int64
}

type Identity struct {
	IdentityData
}

type ArticleData struct {
	Id              *int64
	Excerpt         *string
	Url             *string
	CreatedAt       *time.Time
	Title           *string
	ArticleSourceId *int64
	ThumbnailSrc    *string
	Tags            *Strings
}

type Article struct {
	ArticleData
	ArticleSource *ArticleSource
	Claim         []*Claim
}
