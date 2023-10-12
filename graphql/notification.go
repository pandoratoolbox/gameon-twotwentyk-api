package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_notification = ReflectToFragment(models.NotificationData{})
)

func NewNotification(ctx context.Context, data *models.Notification) error {
	q := `
		mutation CreateNotification {
			notification(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Notification
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
		Notification []models.Notification
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Notification) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Notification[0].Id

	data.Id = &id

	return nil
}

func DeleteNotification(ctx context.Context, id int64) error {
	q := `
		mutation DeleteNotification {
			notification(delete: true, where: { id: { eq: $id } }) {
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
		Notification []models.Notification
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Notification) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateNotification(ctx context.Context, data models.Notification) error {
	q := `
		mutation UpdateNotification {
			notification(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Notification
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
		Notification []models.Notification
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Notification) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetNotification(ctx context.Context, id int64) (models.Notification, error) {
	var data models.Notification

	q := fragment_notification + `
			query GetNotification {
			notification(where: { id: { eq: $id } }) {
				...Notification
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
		Notification []models.Notification
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Notification) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Notification[0]

	return data, nil
}

func ListNotificationByRecipientId(ctx context.Context, id int64) ([]models.Notification, error) {
	var out []models.Notification

	q := fragment_notification + `query ListNotificationByRecipientId(where: { recipient_id: { eq: $id }}) {
						...Notification
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
		Notification []models.Notification
	}{}

	err = json.Unmarshal(res.Data, &ret)
	if err != nil {
		return out, err
	}

	out = ret.Notification

	return out, nil
}
