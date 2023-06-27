package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewMarketplaceListing(ctx context.Context, data *models.MarketplaceListing) error {
	err := graphql.NewMarketplaceListing(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetMarketplaceListing(ctx context.Context, id int64) (models.MarketplaceListing, error) {
	data, err := graphql.GetMarketplaceListing(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteMarketplaceListing(ctx context.Context, id int64) error {
	err := graphql.DeleteMarketplaceListing(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateMarketplaceListing(ctx context.Context, data models.MarketplaceListing) error {
	err := graphql.UpdateMarketplaceListing(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListMarketplaceListingByOwnerId(ctx context.Context, id int64) ([]models.MarketplaceListing, error) {
	data, err := graphql.ListMarketplaceListingByOwnerId(ctx, id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func SearchMarketplaceListings(ctx context.Context, q string) ([]models.MarketplaceListing, error) {
	data, err := graphql.SearchMarketplaceListings(ctx, q)
	if err != nil {
		return data, err
	}

	return data, nil
}
