package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewVerification(ctx context.Context, data *models.Verification) error {
	err := graphql.NewVerification(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetVerification(ctx context.Context, id int64) (models.Verification, error) {
	data, err := graphql.GetVerification(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteVerification(ctx context.Context, id int64) error {
	err := graphql.DeleteVerification(ctx, id)
	if err != nil {
		return err
	}

	return nil
}