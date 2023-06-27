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

type MarketplaceListingData struct {
	CreatedAt       *time.Time
	IsListed        *bool
	NftCollectionId *int64
	NftId           *int64
	Id              *int64
	OwnerId         *int64
	Price           *int64
}

type MarketplaceListing struct {
	MarketplaceListingData
	Owner *User
}

type CelebrityData struct {
	Categories *Strings
	Id         *int64
	Name       *string
	BirthDay   *int64
	BirthMonth *int64
	BirthYear  *int64
}

type Celebrity struct {
	CelebrityData
}

type CardPackSeriesData struct {
	Name        *string
	PctIdentity *int64
	CardAmount  *int64
	Id          *int64
	Quantity    *int64
	PctYear     *int64
	PctMonth    *int64
	PctDay      *int64
	PctEvent    *int64
}

type CardPackSeries struct {
	CardPackSeriesData
	CardPacks []*CardPacks
}

type NftCardIdentityData struct {
	Month         *int64
	Year          *int64
	IsCrafted     *bool
	OwnerId       *int64
	Category      *string
	Rarity        *int64
	Id            *int64
	Day           *int64
	CelebrityName *string
}

type NftCardIdentity struct {
	NftCardIdentityData
	Owner *User
}

type NftCardPredictionData struct {
	Rarity        *int64
	Id            *int64
	IsClaimed     *bool
	Triggers      *Strings
	CelebrityName *string
	OwnerId       *int64
}

type NftCardPrediction struct {
	NftCardPredictionData
	Owner *User
}

type ArticleData struct {
	Url             *string
	CreatedAt       *time.Time
	Excerpt         *string
	ArticleSourceId *int64
	ThumbnailSrc    *string
	Title           *string
	Tags            *Strings
	Id              *int64
}

type Article struct {
	ArticleData
	Claim         []*Claim
	ArticleSource *ArticleSource
}

type ArticleSourceData struct {
	Id   *int64
	Name *string
}

type ArticleSource struct {
	ArticleSourceData
	Article []*Article
}

type NftCardTriggerData struct {
	Id        *int64
	IsCrafted *bool
	OwnerId   *int64
	Trigger   *string
	Rarity    *int64
}

type NftCardTrigger struct {
	NftCardTriggerData
	Owner *User
}

type CardPacksData struct {
	Id               *int64
	CardPackSeriesId *int64
	Cards            *map[string]interface{}
}

type CardPacks struct {
	CardPacksData
	CardPackSeries *CardPackSeries
}

type NftCardCategoryData struct {
	Id        *int64
	Category  *string
	OwnerId   *int64
	IsCrafted *bool
	Rarity    *int64
}

type NftCardCategory struct {
	NftCardCategoryData
	Owner *User
}

type NftCardDayMonthData struct {
	Id        *int64
	Day       *int64
	Month     *int64
	OwnerId   *int64
	Rarity    *int64
	IsCrafted *bool
}

type NftCardDayMonth struct {
	NftCardDayMonthData
	Owner *User
}

type UserData struct {
	Name           *string
	Id             *int64
	Email          *string
	PhoneNumber    *string
	Username       *string
	WalletAddress  *string
	CreatedAt      *time.Time
	Password       *string
	ExternalAuthId *string
	RoleIds        *Ints
}

type User struct {
	UserData
	MarketplaceListing []*MarketplaceListing
	NftCardPrediction  []*NftCardPrediction
	NftCardDayMonth    []*NftCardDayMonth
	NftCardIdentity    []*NftCardIdentity
	NftCardCrafting    []*NftCardCrafting
	NftCardTrigger     []*NftCardTrigger
	NftCardCategory    []*NftCardCategory
	Claim              []*Claim
	NftCardYear        []*NftCardYear
}

type ClaimData struct {
	Id              *int64
	Status          *int64
	CreatedAt       *time.Time
	ClaimerId       *int64
	NftPredictionId *int64
	ArticleId       *int64
}

type Claim struct {
	ClaimData
	Claimer *User
	Article *Article
}

type IdentityData struct {
	Category   *string
	Id         *int64
	Name       *string
	BirthDay   *int64
	BirthMonth *int64
	BirthYear  *int64
}

type Identity struct {
	IdentityData
}

type CategoryData struct {
	Id   *int64
	Name *string
}

type Category struct {
	CategoryData
}

type NftCardYearData struct {
	Rarity    *int64
	Id        *int64
	Year      *int64
	OwnerId   *int64
	IsCrafted *bool
}

type NftCardYear struct {
	NftCardYearData
	Owner *User
}

type NftCardCraftingData struct {
	Id        *int64
	IsCrafted *bool
	OwnerId   *int64
	Rarity    *int64
}

type NftCardCrafting struct {
	NftCardCraftingData
	Owner *User
}

type TriggerData struct {
	Id   *int64
	Name *string
}

type Trigger struct {
	TriggerData
}
