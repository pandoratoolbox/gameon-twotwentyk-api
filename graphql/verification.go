package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	Fragment_verification = ReflectToFragment(models.VerificationData{})
)

func NewVerification(ctx context.Context, data *models.Verification) error {
	q := `
		mutation CreateVerification {
			verification(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Verification
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
		Verification []models.Verification
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Verification) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Verification[0].Id

	data.Id = &id

	return nil
}

func GetVerification(ctx context.Context, id int64) (models.Verification, error) {
	var data models.Verification

	q := Fragment_verification + `
			query GetVerification {
			verification(order_by: {created_at: desc}, where: { user_id: { eq: $id } }) {
				...Verification
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
		Verification []models.Verification
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Verification) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Verification[0]

	return data, nil
}

func DeleteVerification(ctx context.Context, id int64) error {
	q := `
		mutation DeleteVerification {
			verification(delete: true, where: { id: { eq: $id } }) {
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
		Verification []models.Verification
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Verification) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}