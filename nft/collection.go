package nft

const WALLET_KEY_PRIVATE = "wallet_key_private"

// type AggPack struct {
// 	AmountNftCard          int64
// 	AmountNftCardCelebrity int64
// 	AmountNftCardCategory  int64
// 	AmountNftCardDayMonth  int64
// 	AmountNftCardTrigger   int64
// }

// type AggPackRaw struct {
// 	Standard []map[int64]int64
// 	Premium  []map[int64]int64
// 	Elite    []map[int64]int64
// }

// type CollectionConfig struct {
// }

// type CollectionConfigTier struct {
// 	CardAmount   int64
// 	CardsPerPack []CollectionConfigCardsPerPack
// 	Price        int64
// 	Name         string
// }

// type CollectionConfigCardsPerPack struct {
// 	NftTypeId int64
// 	Amount    int64
// }

// type CollectionConfigPack struct {
// 	Tiers []CollectionConfigTier
// }

// func CreateCardCollection(data models.CardCollection, agg_pack AggPack) error {

// 	return nil
// }

// func OpenCardPack(ctx context.Context, card_collection_id int64, card_pack_id int64) error {

// 	collection, err := store.GetCardCollection(ctx, card_collection_id)
// 	if err != nil {
// 		return err
// 	}

// 	card_pack, err := store.GetCardPack(ctx, card_pack_id)
// 	if err != nil {
// 		return err
// 	}

// 	owner, err := store.GetUser(ctx, *card_pack.OwnerId)
// 	if err != nil {
// 		return err
// 	}

// 	owner_address := *owner.WalletAddress

// 	// card_pack_config, err := store.GetCardPackConfig(ctx, card_collection_id, card_pack_id)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	card_pacK_config := models.CardPackConfig{}

// 	for _, core_trigger := range card_pacK_config.Contains.Core.Trigger {
// 		rarity := models.NFT_RARITY_COMMON
// 		triggers, err := store.ListTriggers(ctx)
// 		if err != nil {
// 			return err
// 		}
// 		rnd_trigger := triggers(rand.Intn(len(triggers)))
// 		// trigger_tier := rnd_trigger.Tier
// 		// id := rnd_trigger.Id

// 		trigger_name := rnd_trigger.Name
// 		trigger_tier := rnd_trigger.Tier

// 		is_crafted := false

// 		nft_trigger := models.NftCardTrigger{
// 			NftCardTriggerData: models.NftCardTriggerData{
// 				Trigger:   &trigger_name,
// 				Tier:      &trigger_tier,
// 				IsCrafted: &is_crafted,
// 				Rarity:    &rarity,
// 			},
// 		}

// 		err = store.NewNftCardTrigger(ctx, &nft_trigger)
// 		if err != nil {
// 			return err
// 		}

// 		tokenId := *nft_trigger.Id

// 		err = nft.CreateTriggerCard(owner_address, collectionId, tokenId, tokenURI, nft_trigger)
// 		if err != nil {
// 			return err
// 		}

// 	}

// 	for _, core_year := range card_pacK_config.Contains.Core.Year {
// 		rarity := models.NFT_RARITY_COMMON
// 		celebrities, err := store.ListCelebrity(ctx)
// 		if err != nil {
// 			return err
// 		}
// 		rnd_celebrity := celebrities(rand.Intn(len(celebrities)))
// 		year := rnd_celebrity.BirthYear
// 		is_crafted := false

// 		nft_day_year := models.NftCardYear{
// 			NftCardYearData: models.NftCardYearData{
// 				Year:      &year,
// 				IsCrafted: &is_crafted,
// 				Rarity:    &rarity,
// 			},
// 		}

// 		err = store.NewNftCardYear(ctx, &nft_day_year)
// 		if err != nil {
// 			return err
// 		}

// 		tokenId := *nft_day_year.Id

// 		err = nft.CreateYearCard(owner_address, collectionId, tokenId, tokenURI, nft_day_year)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for _, core_day_month := range card_pacK_config.Contains.Core.DayMonth {
// 		rarity := models.NFT_RARITY_COMMON
// 		celebrities, err := store.ListCelebrity(ctx)
// 		if err != nil {
// 			return err
// 		}
// 		rnd_celebrity := celebrities(rand.Intn(len(celebrities)))
// 		day := rnd_celebrity.BirthDay
// 		month := rnd_celebrity.BirthMonth
// 		is_crafted := false

// 		nft_day_month := models.NftCardDayMonth{
// 			NftCardDayMonthData: models.NftCardDayMonthData{
// 				Day:       &day,
// 				Month:     &month,
// 				IsCrafted: &is_crafted,
// 				Rarity:    &rarity,
// 			},
// 		}

// 		err = store.NewNftCardDayMonth(ctx, &nft_day_month)
// 		if err != nil {
// 			return err
// 		}

// 		tokenId := *nft_day_month.Id

// 		err = nft.CreateDayMonthCard(owner_address, collectionId, tokenId, tokenURI, nft_day_month)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	for _, core_category := range card_pacK_config.Contains.Core.Category {

// 		//get random category
// 		rarity := models.NFT_RARITY_COMMON
// 		celebrities, err := store.ListCelebrity(ctx)
// 		if err != nil {
// 			return err
// 		}
// 		rnd_celebrity := celebrities(rand.Intn(len(celebrities)))
// 		category_name := rnd_celebrity.Category
// 		is_crafted := false

// 		nft_category := models.NftCardCategory{
// 			NftCardCategoryData: models.NftCardCategoryData{
// 				Name:      &category_name,
// 				Rarity:    &rarity,
// 				IsCrafted: &is_crafted,
// 				Category:  &category_name,
// 			},
// 		}

// 		err = store.NewNftCardCategory(ctx, &nft_category)
// 		if err != nil {
// 			return err
// 		}

// 		tokenId := *nft_category.Id

// 		err = nft.CreateCategoryCard(owner_address, collectionId, tokenId, tokenURI, nft_category)
// 		if err != nil {
// 			return err
// 		}
// 		// mint core
// 	}

// 	return nil
// }
