package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewCelebrity(ctx context.Context, data *models.Celebrity) error {
	err := graphql.NewCelebrity(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetCelebrity(ctx context.Context, id int64) (models.Celebrity, error) {
	data, err := graphql.GetCelebrity(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteCelebrity(ctx context.Context, id int64) error {
	err := graphql.DeleteCelebrity(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCelebrity(ctx context.Context, data models.Celebrity) error {
	err := graphql.UpdateCelebrity(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListCelebrity(ctx context.Context) ([]models.Celebrity, error) {
	out, err := graphql.ListCelebrity(ctx)
	if err != nil {
		return out, err
	}

	return out, nil
}
