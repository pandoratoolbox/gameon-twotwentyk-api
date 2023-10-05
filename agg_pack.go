package main

import (
	"errors"
	"gameon-twotwentyk-api/models"
)

type AggPack struct {
	Tier [][]map[int64]interface{}

	// 3 x tier -> 12 x card type -> array of card amount values for each card pack in collection
}

// func DecodeAggPack(collection models.CardCollection, path string) (map[string]map[int64]models.CardPackConfig, error) {

// 	b, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// fmt.Println(string(b))

// 	var agg_pack AggPack
// 	err = json.Unmarshal(b, &agg_pack.Tier)
// 	if err != nil {
// 		return nil, err
// 	}
// 	cards := make(map[string]map[int64]models.CardPackConfig)

// 	for s, series := range collection.CardSeries {
// 		cards_per_pack := *series.CardsPerPack
// 		index := int64(0)

// 		for a := 0; a < int(*series.CardPackQuantity); a++ {
// 			card := models.CardPackConfig{}

// 			for i := 0; i < int(cards_per_pack); i++ {

// 				for cti := 0; cti < (4 * len(rarities)); cti++ {
// 					//for i := 0; i < 4; i++ //loop over card types for each rarity
// 					v := agg_pack.Tier[s][cti][index]

// 					if fmt.Sprintf("%T", v) == "float64" {
// 						if v.(float64) == 0 {
// 							continue
// 						}
// 					}

// 					var card models.CardPackConfig
// 					c, ok := cards[*series.Name][int64(a)]
// 					if ok {
// 						card = c
// 					}

// 					card.CardPackId = int64(a) + 1
// 					card.Tier = *series.Name

// 					switch cti {
// 					case 0:
// 						card.Contains.Rare.Year++
// 					case 1:
// 						card.Contains.Rare.DayMonth++
// 					case 2:
// 						card.Contains.Rare.Category++
// 					case 3:
// 						if fmt.Sprintf("%T", v) == "string" {
// 							t_strs := strings.Split(strings.TrimSpace(v.(string)), "_")
// 							t_tier := ""
// 							if t_strs[1] == "Minor" {
// 								t_tier = "Minor_" + t_strs[2]
// 							}

// 							if t_strs[1] == "Major" {
// 								t_tier = "Major"
// 							}

// 							t_id, err := strconv.ParseInt(t_strs[len(t_strs)-1], 10, 64)
// 							if err != nil {
// 								return nil, err
// 							}

// 							card.Contains.Rare.Trigger = append(card.Contains.Rare.Trigger, struct {
// 								Name string
// 								Tier string
// 								Id   int64
// 							}{
// 								Tier: t_tier,
// 								Id:   t_id,
// 							})

// 						}
// 					case 4:
// 						card.Contains.Uncommon.Year++
// 					case 5:
// 						card.Contains.Uncommon.DayMonth++
// 					case 6:
// 						card.Contains.Uncommon.Category++
// 					case 7:
// 						if fmt.Sprintf("%T", v) == "string" {
// 							t_strs := strings.Split(strings.TrimSpace(v.(string)), "_")
// 							t_tier := ""
// 							if t_strs[1] == "Minor" {
// 								t_tier = "Minor_" + t_strs[2]
// 							}

// 							if t_strs[1] == "Major" {
// 								t_tier = "Major"
// 							}

// 							t_id := t_strs[len(t_strs)-1]

// 							card.Contains.Uncommon.Trigger.Id, err = strconv.ParseInt(t_id, 10, 64)
// 							if err != nil {
// 								return nil, err
// 							}

// 							card.Contains.Uncommon.Trigger.Tier = t_tier
// 						}
// 					case 8:
// 						card.Contains.Core.Year++
// 					case 9:
// 						card.Contains.Core.DayMonth++
// 					case 10:
// 						card.Contains.Core.Category++
// 					case 11:
// 						if fmt.Sprintf("%T", v) == "string" {
// 							t_strs := strings.Split(strings.TrimSpace(v.(string)), "_")
// 							t_tier := ""
// 							if t_strs[1] == "Minor" {
// 								t_tier = "Minor_" + t_strs[2]
// 							}

// 							if t_strs[1] == "Major" {
// 								t_tier = "Major"
// 							}

// 							t_id := t_strs[len(t_strs)-1]

// 							card.Contains.Core.Trigger.Id, err = strconv.ParseInt(t_id, 10, 64)
// 							if err != nil {
// 								return nil, err
// 							}

// 							card.Contains.Core.Trigger.Tier = t_tier

// 							amount, err := strconv.ParseInt(t_strs[0], 10, 64)
// 							if err != nil {
// 								return nil, err
// 							}

// 							card.Contains.Core.Trigger.Amount = amount

// 						}
// 					}
// 				}

// 				index++
// 			}

// 			cards[*series.Name][int64(a)] = card

// 		}
// 	}
// 	js, err := json.Marshal(cards)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// fmt.Println(string(js))

// 	err = ioutil.WriteFile("agg_out.json", js, 0644)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return cards, nil
// }

func GenerateAggPack(collection models.CardCollection) error {
	if *collection.Status != 3 {
		return errors.New("Status must be 3 to generate agg pack")
	}

	return nil
}
