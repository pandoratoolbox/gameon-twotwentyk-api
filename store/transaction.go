package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewTransaction(ctx context.Context, data *models.Transaction) error {
	err := graphql.NewTransaction(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetTransaction(ctx context.Context, id int64) (models.Transaction, error) {
	data, err := graphql.GetTransaction(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteTransaction(ctx context.Context, id int64) error {
	err := graphql.DeleteTransaction(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTransaction(ctx context.Context, data models.Transaction) error {
	err := graphql.UpdateTransaction(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
