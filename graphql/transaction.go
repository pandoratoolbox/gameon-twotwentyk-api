package graphql

import (
	"context"
	"errors"
	"gameon-twotwentyk-api/models"

	"github.com/pandoratoolbox/json"
)

var (
	fragment_transaction = ReflectToFragment(models.TransactionData{})
)

func NewTransaction(ctx context.Context, data *models.Transaction) error {
	q := `
		mutation CreateTransaction {
			transaction(insert: $data) {
				id
			}
		}
		`

	input := struct {
		Data models.Transaction
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
		Transaction []models.Transaction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Transaction) < 1 {
		return errors.New("Unable to insert object")
	}

	id := *out.Transaction[0].Id

	data.Id = &id

	return nil
}

func DeleteTransaction(ctx context.Context, id int64) error {
	q := `
		mutation DeleteTransaction {
			transaction(delete: true, where: { id: { eq: $id } }) {
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
		Transaction []models.Transaction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Transaction) < 1 {
		return errors.New("Unable to delete object")
	}

	return nil
}

func UpdateTransaction(ctx context.Context, data models.Transaction) error {
	q := `
		mutation UpdateTransaction {
			transaction(where: { id: { eq: $id } }, update: $data) {
				id
			}
		}
		`

	input := struct {
		Id   int64
		Data models.Transaction
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
		Transaction []models.Transaction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return err
	}

	if len(out.Transaction) < 1 {
		return errors.New("Unable to update object")
	}

	return nil
}

func GetTransaction(ctx context.Context, id int64) (models.Transaction, error) {
	var data models.Transaction

	q := fragment_transaction + `
			query GetTransaction {
			transaction(where: { id: { eq: $id } }) {
				...Transaction
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
		Transaction []models.Transaction
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return data, err
	}

	if len(out.Transaction) < 1 {
		return data, errors.New("Unable to retrieve object")
	}

	data = out.Transaction[0]

	return data, nil
}
