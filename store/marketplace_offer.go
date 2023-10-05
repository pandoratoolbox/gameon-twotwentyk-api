package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewMarketplaceOffer(ctx context.Context, data *models.MarketplaceOffer) error {
	err := graphql.NewMarketplaceOffer(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetMarketplaceOffer(ctx context.Context, id int64) (models.MarketplaceOffer, error) {
	data, err := graphql.GetMarketplaceOffer(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteMarketplaceOffer(ctx context.Context, id int64) error {
	err := graphql.DeleteMarketplaceOffer(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateMarketplaceOffer(ctx context.Context, data models.MarketplaceOffer) error {
	err := graphql.UpdateMarketplaceOffer(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListMarketplaceOfferByBuyerId(ctx context.Context, id int64) ([]models.MarketplaceOffer, error) {
	data, err := graphql.ListMarketplaceOfferByBuyerId(ctx, id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func ListMarketplaceOfferByMarketplaceListingId(ctx context.Context, id int64) ([]models.MarketplaceOffer, error) {
	data, err := graphql.ListMarketplaceOfferByMarketplaceListingId(ctx, id)
	if err != nil {
		return data, err
	}

	return data, nil
}
