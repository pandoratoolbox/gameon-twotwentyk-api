package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_prediction = ReflectToFragment(models.NftCardPredictionData{})
)

func NewNftCardPrediction(ctx context.Context, data *models.NftCardPrediction) error {
	q := `
		mutation CreateNftCardPrediction {
			nft_card_prediction(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardPrediction
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
		NftCardPrediction []models.NftCardPrediction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardPrediction) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardPrediction[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardPrediction(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardPrediction {
			nft_card_prediction(where: { id: { eq: $id } }) {
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
		NftCardPrediction []models.NftCardPrediction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardPrediction) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardPrediction(ctx context.Context, data models.NftCardPrediction) error {
	q := `
		mutation UpdateNftCardPrediction {
			nft_card_prediction(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardPrediction
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
		NftCardPrediction []models.NftCardPrediction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardPrediction) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardPrediction(ctx context.Context, id int64) (models.NftCardPrediction, error) {
	var data models.NftCardPrediction

	q := fragment_nft_card_prediction + `
			query GetNftCardPrediction {
			nft_card_prediction(where: { id: { eq: $id } }) {
				...NftCardPrediction
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
		NftCardPrediction []models.NftCardPrediction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardPrediction) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardPrediction[0]

	return data, nil
}

func ListNftCardPredictionByOwnerId(ctx context.Context, id int64) ([]models.NftCardPrediction, error) {
	var out []models.NftCardPrediction

	q := fragment_nft_card_prediction + `query ListNftCardPredictionByOwnerId(where: { owner_id: { eq: $id }}) {
						...NftCardPrediction
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
		NftCardPrediction []models.NftCardPrediction
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	if len(ret.NftCardPrediction) < 1 {
		return out, errors.New("Object not found")
	}

	out = ret.NftCardPrediction

	return out, nil
}
