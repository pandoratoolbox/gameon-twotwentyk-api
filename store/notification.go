package store

import (
	"context"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/models"
)

func NewNotification(ctx context.Context, data *models.Notification) error {
	err := graphql.NewNotification(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func GetNotification(ctx context.Context, id int64) (models.Notification, error) {
	data, err := graphql.GetNotification(ctx, id)

	if err != nil {
		return data, err
	}

	return data, nil
}

func DeleteNotification(ctx context.Context, id int64) error {
	err := graphql.DeleteNotification(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func UpdateNotification(ctx context.Context, data models.Notification) error {
	err := graphql.UpdateNotification(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func ListNotificationByRecipientId(ctx context.Context, id int64) ([]models.Notification, error) {
	data, err := graphql.ListNotificationByRecipientId(ctx, id)
	if err != nil {
		return data, err
	}

	return data, nil
}
