package graphql

import (
	"context"
	"encoding/json"
	"gameon-twotwentyk-api/models"
)

var fragment_trigger = ReflectToFragment(models.TriggerData{})

func ListTrigger(ctx context.Context) ([]models.Trigger, error) {
	out := struct {
		Trigger []models.Trigger
	}{}

	q := fragment_trigger + `query ListTrigger {
		trigger() {
			...Trigger
		}
	}`

	res, err := Graph.GraphQL(ctx, q, nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res.Data, &out)
	if err != nil {
		return nil, err
	}

	return out.Trigger, nil
}
