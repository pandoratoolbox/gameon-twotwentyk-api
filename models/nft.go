package models

import "time"

type NftCardDay struct {
	Id  int64
	Day int
}

type NftCardMonth struct {
	Id    int64
	Month int
}

type NftCardYear struct {
	Id   int64
	Year int
}

type NftCardIdentity struct {
	Id int64
	// CelebrityId int64
	CelebrityName      string
	CelebrityBirthDate time.Time
}

type NftCardPrediction struct {
	Id            int64
	EventName     string
	CelebrityName string
}

type NftCardTrigger struct {
	Id        int64
	EventName string
}
