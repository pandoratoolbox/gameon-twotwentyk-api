package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_marketplace_offer = ReflectToFragment(models.MarketplaceOfferData{})
)

func NewMarketplaceOffer(ctx context.Context, data *models.MarketplaceOffer) error {
	q := `
		mutation CreateMarketplaceOffer {
			marketplace_offer(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.MarketplaceOffer
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
		MarketplaceOffer []models.MarketplaceOffer
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.MarketplaceOffer) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.MarketplaceOffer[0].Id

	data.Id = &id

	return nil
}

func DeleteMarketplaceOffer(ctx context.Context, id int64) error {
	q := `
		mutation DeleteMarketplaceOffer {
			marketplace_offer(delete: true, where: { id: { eq: $id } }) {
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
		MarketplaceOffer []models.MarketplaceOffer
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.MarketplaceOffer) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateMarketplaceOffer(ctx context.Context, data models.MarketplaceOffer) error {
	q := `
		mutation UpdateMarketplaceOffer {
			marketplace_offer(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.MarketplaceOffer
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
		MarketplaceOffer []models.MarketplaceOffer
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.MarketplaceOffer) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetMarketplaceOffer(ctx context.Context, id int64) (models.MarketplaceOffer, error) {
	var data models.MarketplaceOffer

	q := fragment_marketplace_offer + `
			query GetMarketplaceOffer {
			marketplace_offer(where: { id: { eq: $id } }) {
				...MarketplaceOffer
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
		MarketplaceOffer []models.MarketplaceOffer
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.MarketplaceOffer) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.MarketplaceOffer[0]

	return data, nil
}

func ListMarketplaceOfferByBuyerId(ctx context.Context, id int64) ([]models.MarketplaceOffer, error) {
	var out []models.MarketplaceOffer

	q := fragment_marketplace_offer + `query ListMarketplaceOfferByBuyerId {
		marketplace_offer(where: { buyer_id: { eq: $id }}) {
						...MarketplaceOffer
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
		MarketplaceOffer []models.MarketplaceOffer
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.MarketplaceOffer

	return out, nil
}

func ListMarketplaceOfferByMarketplaceListingId(ctx context.Context, id int64) ([]models.MarketplaceOffer, error) {
	var out []models.MarketplaceOffer

	q := fragment_marketplace_offer + `query ListMarketplaceOfferByMarketplaceListingId(where: { marketplace_listing_id: { eq: $id }}) {
						...MarketplaceOffer
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
		MarketplaceOffer []models.MarketplaceOffer
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	if len(ret.MarketplaceOffer) < 1 {
		return out, errors.New("Object not found")
	}

	out = ret.MarketplaceOffer

	return out, nil
}
