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

type ArticleData struct {
	ThumbnailSrc    *string
	Title           *string
	Tags            *Strings
	ArticleSourceId *int64
	Url             *string
	CreatedAt       *time.Time
	Id              *int64
	Excerpt         *string
}

type Article struct {
	ArticleData
	ArticleSource *ArticleSource
	Claim         []*Claim
}

type ArticleSourceData struct {
	Id   *int64
	Name *string
}

type ArticleSource struct {
	ArticleSourceData
	Article []*Article
}

type CardPackSeriesData struct {
	Id          *int64
	Name        *string
	Quantity    *int64
	PctYear     *int64
	PctIdentity *int64
	PctMonth    *int64
	PctDay      *int64
	PctEvent    *int64
	CardAmount  *int64
}

type CardPackSeries struct {
	CardPackSeriesData
	CardPacks []*CardPacks
}

type CardPacksData struct {
	CardPackSeriesId *int64
	Cards            *map[string]interface{}
	Id               *int64
}

type CardPacks struct {
	CardPacksData
	CardPackSeries *CardPackSeries
}

type UserData struct {
	CreatedAt      *time.Time
	Password       *string
	ExternalAuthId *string
	Id             *int64
	Username       *string
	Email          *string
	WalletAddress  *string
	RoleIds        *Ints
}

type User struct {
	UserData
	Claim              []*Claim
	MarketplaceListing []*MarketplaceListing
}

type MarketplaceListingData struct {
	Id              *int64
	OwnerId         *int64
	Price           *int64
	CreatedAt       *time.Time
	IsListed        *bool
	NftCollectionId *int64
	NftId           *int64
}

type MarketplaceListing struct {
	MarketplaceListingData
	Owner *User
}

type CelebrityData struct {
	Id        *int64
	Name      *string
	Birthdate *time.Time
}

type Celebrity struct {
	CelebrityData
}
