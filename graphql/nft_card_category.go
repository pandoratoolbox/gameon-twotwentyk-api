package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_nft_card_category = ReflectToFragment(models.NftCardCategoryData{})
)

func NewNftCardCategory(ctx context.Context, data *models.NftCardCategory) error {
	q := `
		mutation CreateNftCardCategory {
			nft_card_category(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.NftCardCategory
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
		NftCardCategory []models.NftCardCategory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardCategory) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.NftCardCategory[0].Id

	data.Id = &id

	return nil
}

func DeleteNftCardCategory(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNftCardCategory {
			nft_card_category(where: { id: { eq: $id } }) {
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
		NftCardCategory []models.NftCardCategory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardCategory) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNftCardCategory(ctx context.Context, data models.NftCardCategory) error {
	q := `
		mutation UpdateNftCardCategory {
			nft_card_category(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.NftCardCategory
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
		NftCardCategory []models.NftCardCategory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.NftCardCategory) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNftCardCategory(ctx context.Context, id int64) (models.NftCardCategory, error) {
	var data models.NftCardCategory

	q := fragment_nft_card_category + `
			query GetNftCardCategory {
			nft_card_category(where: { id: { eq: $id } }) {
				...NftCardCategory
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
		NftCardCategory []models.NftCardCategory
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.NftCardCategory) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.NftCardCategory[0]

	return data, nil
}

func ListNftCardCategoryByOwnerId(ctx context.Context, id int64) ([]models.NftCardCategory, error) {
	var out []models.NftCardCategory

	q := fragment_nft_card_category + `query ListNftCardCategoryByOwnerId(where: { owner_id: { eq: $id }}) {
						...NftCardCategory
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
		NftCardCategory []models.NftCardCategory
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	if len(ret.NftCardCategory) < 1 {
		return out, errors.New("Object not found")
	}

	out = ret.NftCardCategory

	return out, nil
}
