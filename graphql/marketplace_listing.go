package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_marketplace_listing = ReflectToFragment(models.MarketplaceListingData{})
)

func NewMarketplaceListing(ctx context.Context, data *models.MarketplaceListing) error {
	q := `
		mutation CreateMarketplaceListing {
			marketplace_listing(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.MarketplaceListing
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
		MarketplaceListing []models.MarketplaceListing
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.MarketplaceListing) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.MarketplaceListing[0].Id

	data.Id = &id

	return nil
}

func DeleteMarketplaceListing(ctx context.Context, id int64) error {
	q := `
		mutation DeleteMarketplaceListing {
			marketplace_listing(where: { id: { eq: $id } }) {
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
		MarketplaceListing []models.MarketplaceListing
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.MarketplaceListing) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateMarketplaceListing(ctx context.Context, data models.MarketplaceListing) error {
	q := `
		mutation UpdateMarketplaceListing {
			marketplace_listing(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.MarketplaceListing
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
		MarketplaceListing []models.MarketplaceListing
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.MarketplaceListing) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetMarketplaceListing(ctx context.Context, id int64) (models.MarketplaceListing, error) {
	var data models.MarketplaceListing

	q := fragment_marketplace_listing + `
			query GetMarketplaceListing {
			marketplace_listing(where: { id: { eq: $id } }) {
				...MarketplaceListing
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
		MarketplaceListing []models.MarketplaceListing
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.MarketplaceListing) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.MarketplaceListing[0]

	return data, nil
}

func ListMarketplaceListingByOwnerId(ctx context.Context, id int64) ([]models.MarketplaceListing, error) {
	var out []models.MarketplaceListing

	q := fragment_marketplace_listing + `query ListMarketplaceListingByOwnerId {
		marketplace_listing(where: { owner_id: { eq: $id }}) {
						...MarketplaceListing
		}
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
		MarketplaceListing []models.MarketplaceListing
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	// if len(ret.MarketplaceListing) < 1 {
	// 	return out, errors.New("Object not found")
	// }

	out = ret.MarketplaceListing

	return out, nil
}

func SearchMarketplaceListings(ctx context.Context, q string) ([]models.MarketplaceListing, error) {
	var out []models.MarketplaceListing

	gq := fragment_marketplace_listing + `query SearchMarketplaceListings {
		marketplace_listing() {
						...MarketplaceListing
					}
				}`

	// 	gq := fragment_article + `
	// 	query ListArticles {
	// 	article() {
	// 		...Article
	// 	}
	// }
	// `

	// input := struct {
	// 	Q string `json:"q"`
	// }{
	// 	Q: q,
	// }

	// js, err := json.Marshal(input)
	// if err != nil {
	// 	return out, err
	// }

	res, err := Graph.GraphQL(ctx, gq, nil, nil)
	if err != nil {
		return out, err
	}

	ret := struct {
		MarketplaceListing []models.MarketplaceListing
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	// if len(ret.MarketplaceListing) < 1 {
	// 	return out, errors.New("Object not found")
	// }

	out = ret.MarketplaceListing

	return out, nil
}
