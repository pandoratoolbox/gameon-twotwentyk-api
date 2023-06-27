package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewCategory(ctx context.Context, data *models.Category) error {
	err := graphql.NewCategory(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetCategory(ctx context.Context, id int64) (models.Category, error) {
	data, err := graphql.GetCategory(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteCategory(ctx context.Context, id int64) error {
	err := graphql.DeleteCategory(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCategory(ctx context.Context, data models.Category) error {
	err := graphql.UpdateCategory(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListCategory(ctx context.Context) ([]models.Category, error) {
	data, err := graphql.ListCategory(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
