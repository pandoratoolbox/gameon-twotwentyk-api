package nft

// func TransferNFTViaVenly(from string, to string, token string, amount int64) {

// }

// func TransferNFTViaMatic(from string, to string, token string, amount int64) {

// }

// func MintCardPack(tier int64, to string) {

// }

// func MintCategoryCard(category_id int64, to string) {

// }

// func SubmitCollectionConfig(ctx context.Context, collection_id int64) error {
// 	collection, err := store.GetCardCollection(ctx, collection_id)
// 	if err != nil {
// 		return err
// 	}

// 	if *collection.Status != models.CARD_COLLECTION_STATUS_PENDING_CONFIG {
// 		return errors.New("Card collection is not pending config")
// 	}

// 	err = GenerateAggPack(collection)

// }

// func GenerateAggPack(collection models.CardCollection) error {
// 	// agg_pack := AggPack{
// 	// 	AmountNftCard:          0,
// 	// 	AmountNftCardCelebrity: 0,
// 	// 	AmountNftCardCategory:  0,
// 	// 	AmountNftCardTrigger:   0,
// 	// }

// 	return nil

// }

// // func EncodeAggPackConfig(collection models.CardCollection) (AggPackConfig, error) {
// // 	return out, nil
// // }

// func MintCollection(ctx context.Context, collection_id int64) error {

// 	card_collection, err := store.GetCardCollection(ctx, collection_id)
// 	if err != nil {
// 		return err
// 	}

// 	if *card_collection.Status != models.CARD_COLLECTION_STATUS_READY {
// 		return errors.New("Card collection is not ready")
// 	}

// 	var card_pack_configs map[string]map[int64]map[int64]models.CardPackConfig //tier - rarity- id - card pack config
// 	//create card pack configs

// 	DecodeAggPack()

// 	// err = CreateCardCollection(*card_collection, agg_pack)

// 	if err != nil {
// 		return err
// 	}

// 	live := models.CARD_COLLECTION_STATUS_LIVE
// 	upd := models.CardCollection{
// 		CardCollectionData: models.CardCollectionData{
// 			Id:     &collection_id,
// 			Status: &live,
// 		},
// 	}

// 	store.UpdateCardCollection(context.Background(), collection_id, &upd)

// 	store.RefreshCategoryMap(context.Background())
// 	store.RefreshCelebrityMap(context.Background())
// 	store.RefreshTriggerMap(context.Background())
// }

// func CreateCardPack(card_pack_id int64) {

// }

// func OpenCardPack(card_pack_id int64) {

// }

// func GetRandomCelebrity(collection_id int64) {

// }

// func GetRandomTrigger(collection_id int64) {
// }

// type AggPack struct {
// 	Tier [][]map[int64]interface{}

// 	// 3 x tier -> 12 x card type -> array of card amount values for each card pack in collection
// }

// func DecodeAggPack(collection models.CardCollection, path string) (map[string]map[int64]models.CardPackConfig, error) {

// 	if collection.AggPackPath == nil {
// 		return nil, errors.New("Agg pack path is nil")
// 	}

// 	b, err := ioutil.ReadFile(*collection.AggPackPath)
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
// 		rarity := int64(0)
// 		// card_type_index := 0
// 		for a := 0; a < int(*series.CardPackQuantity); a++ {
// 			card := models.CardPackConfig{}

// 			for i := 0; i < int(cards_per_pack); i++ {

// 				for cti := 0; cti < (4 * len(*collection.Rarities)); cti++ {
// 					if cti == 4 {
// 						cti = 0
// 						rarity++
// 					}
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
// 						prev := card.Contains[rarity]
// 						prev.Year++
// 						card.Contains[rarity] = prev
// 					case 1:
// 						prev := card.Contains[rarity]
// 						prev.DayMonth++
// 						card.Contains[rarity] = prev
// 					case 2:
// 						prev := card.Contains[rarity]
// 						prev.Category++
// 						card.Contains[rarity] = prev
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

// 							prev := card.Contains[rarity]

// 							prev.Trigger = append(prev.Trigger, struct {
// 								Name string
// 								Tier string
// 								Id   int64
// 							}{
// 								Tier: t_tier,
// 								Id:   t_id,
// 							})

// 							card.Contains[rarity] = prev
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
