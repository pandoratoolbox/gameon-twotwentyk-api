package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_day_month = ReflectToFragment(models.NftCardDayMonthData{})
)

func NewNftCardDayMonth(ctx context.Context, data *models.NftCardDayMonth) error {
	q := `
		mutation CreateNftCardDayMonth {
			nft_card_day_month(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardDayMonth
	}{
		Data: *data,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		NftCardDayMonth []models.NftCardDayMonth
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardDayMonth) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardDayMonth[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardDayMonth(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardDayMonth {
			nft_card_day_month(where: { id: { eq: $id } }) {
				id
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		NftCardDayMonth []models.NftCardDayMonth
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardDayMonth) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardDayMonth(ctx context.Context, data models.NftCardDayMonth) error {
	q := `
		mutation UpdateNftCardDayMonth {
			nft_card_day_month(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardDayMonth
	}{
		Id: *data.Id,
	}

	data.Id = nil
	input.Data = data

	js, err := json.Marshal(input)
	if err != nil {
		return err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return err
	}

	var out struct {
		NftCardDayMonth []models.NftCardDayMonth
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardDayMonth) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardDayMonth(ctx context.Context, id int64) (models.NftCardDayMonth, error) {
	var data models.NftCardDayMonth

	q := fragment_nft_card_day_month + `
			query GetNftCardDayMonth {
			nft_card_day_month(where: { id: { eq: $id } }) {
				...NftCardDayMonth
			}
		}
		`

	input := struct {
		Id int64
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return data, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return data, err
	}

	var out struct {
		NftCardDayMonth []models.NftCardDayMonth
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardDayMonth) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardDayMonth[0]

	return data, nil
}

func ListNftCardDayMonthByOwnerId(ctx context.Context, id int64) ([]models.NftCardDayMonth, error) {
	var out []models.NftCardDayMonth

	q := fragment_nft_card_day_month + `query ListNftCardDayMonthByOwnerId(where: { owner_id: { eq: $id }}) {
						...NftCardDayMonth
					}`

	input := struct {
		Id int64 `json:"id"`
	}{
		Id: id,
	}

	js, err := json.Marshal(input)
	if err != nil {
		return out, err
	}

	res, err := Graph.GraphQL(ctx, q, js, nil)
	if err != nil {
		return out, err
	}

	ret := struct {
		NftCardDayMonth []models.NftCardDayMonth
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	if len(ret.NftCardDayMonth) < 1 {
		return out, errors.New("Object not found")
	}

	out = ret.NftCardDayMonth

	return out, nil
}
