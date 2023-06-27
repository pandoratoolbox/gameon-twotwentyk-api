package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func ListTrigger(ctx context.Context) ([]models.Trigger, error) {
	data, err := graphql.ListTrigger(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
