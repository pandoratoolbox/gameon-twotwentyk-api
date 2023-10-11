package nft

type AggPackConfigRarities struct {
	ProbabilityOfCardRarities struct {
		StandardCard struct {
			CoreDayMonth     float64 `json:"Core Day & Month"`
			CoreYear         float64 `json:"Core Year"`
			CoreCategory     float64 `json:"Core Category"`
			CoreTrigger      float64 `json:"Core Trigger"`
			UncommonDayMonth float64 `json:"Uncommon Day&Month"`
			UncommonYear     float64 `json:"Uncommon Year"`
			UncommonCategory float64 `json:"Uncommon Category"`
			UncommonTrigger  float64 `json:"Uncommon Trigger"`
			RareDayMonth     int     `json:"Rare Day&Month"`
			RareYear         int     `json:"Rare Year"`
			RareCategory     int     `json:"Rare Category"`
			RareTrigger      int     `json:"Rare Trigger"`
		} `json:"standard card"`
		PremiumCard struct {
			CoreDayMonth     float64 `json:"Core Day & Month"`
			CoreYear         float64 `json:"Core Year"`
			CoreCategory     float64 `json:"Core Category"`
			CoreTrigger      float64 `json:"Core Trigger"`
			UncommonDayMonth float64 `json:"Uncommon Day&Month"`
			UncommonYear     float64 `json:"Uncommon Year"`
			UncommonCategory float64 `json:"Uncommon Category"`
			UncommonTrigger  float64 `json:"Uncommon Trigger"`
			RareDayMonth     float64 `json:"Rare Day&Month"`
			RareYear         float64 `json:"Rare Year"`
			RareCategory     float64 `json:"Rare Category"`
			RareTrigger      float64 `json:"Rare Trigger"`
		} `json:"premium card"`
		EliteCard struct {
			CoreDayMonth     float64 `json:"Core Day & Month"`
			CoreYear         float64 `json:"Core Year"`
			CoreCategory     float64 `json:"Core Category"`
			CoreTrigger      float64 `json:"Core Trigger"`
			UncommonDayMonth float64 `json:"Uncommon Day&Month"`
			UncommonYear     float64 `json:"Uncommon Year"`
			UncommonCategory float64 `json:"Uncommon Category"`
			UncommonTrigger  float64 `json:"Uncommon Trigger"`
			RareDayMonth     float64 `json:"Rare Day&Month"`
			RareYear         float64 `json:"Rare Year"`
			RareCategory     float64 `json:"Rare Category"`
			RareTrigger      float64 `json:"Rare Trigger"`
		} `json:"elite card"`
	} `json:"Probability of card rarities"`
	GuaranteesPerPack struct {
		StandardPack struct {
			Core struct {
				CoreDayMonth int `json:"Core Day & Month"`
				CoreYear     int `json:"Core Year"`
				CoreCategory int `json:"Core Category"`
				CoreTrigger  int `json:"Core Trigger"`
			} `json:"Core"`
			Uncommon struct {
				UncommonDayMonth int `json:"Uncommon Day & Month"`
				UncommonYear     int `json:"Uncommon Year"`
				UncommonCategory int `json:"Uncommon Category"`
				UncommonTrigger  int `json:"Uncommon Trigger"`
			} `json:"Uncommon"`
			Rare struct {
				RareDayMonth int `json:"Rare Day & Month"`
				RareYear     int `json:"Rare Year"`
				RareCategory int `json:"Rare Category"`
				RareTrigger  int `json:"Rare Trigger"`
			} `json:"Rare"`
		} `json:"standard pack"`
		PremiumPack struct {
			Core struct {
				CoreDayMonth int `json:"Core Day & Month"`
				CoreYear     int `json:"Core Year"`
				CoreCategory int `json:"Core Category"`
				CoreTrigger  int `json:"Core Trigger"`
			} `json:"Core"`
			Uncommon struct {
				UncommonDayMonth int `json:"Uncommon Day & Month"`
				UncommonYear     int `json:"Uncommon Year"`
				UncommonCategory int `json:"Uncommon Category"`
				UncommonTrigger  int `json:"Uncommon Trigger"`
			} `json:"Uncommon"`
			Rare struct {
				RareDayMonth int `json:"Rare Day & Month"`
				RareYear     int `json:"Rare Year"`
				RareCategory int `json:"Rare Category"`
				RareTrigger  int `json:"Rare Trigger"`
			} `json:"Rare"`
		} `json:"premium pack"`
		ElitePack struct {
			Core struct {
				CoreDayMonth int `json:"Core Day & Month"`
				CoreYear     int `json:"Core Year"`
				CoreCategory int `json:"Core Category"`
				CoreTrigger  int `json:"Core Trigger"`
			} `json:"Core"`
			Uncommon struct {
				UncommonDayMonth int `json:"Uncommon Day & Month"`
				UncommonYear     int `json:"Uncommon Year"`
				UncommonCategory int `json:"Uncommon Category"`
				UncommonTrigger  int `json:"Uncommon Trigger"`
			} `json:"Uncommon"`
			Rare struct {
				RareDayMonth int `json:"Rare Day & Month"`
				RareYear     int `json:"Rare Year"`
				RareCategory int `json:"Rare Category"`
				RareTrigger  int `json:"Rare Trigger"`
			} `json:"Rare"`
		} `json:"elite pack"`
	} `json:"Guarantees per pack"`
	TriggerDetail struct {
		StandardPack struct {
			Core struct {
				Major1       float64 `json:"Major_1 "`
				Major2       float64 `json:"Major_2 "`
				Major3       float64 `json:"Major_3 "`
				Major4       float64 `json:"Major_4 "`
				MinorTier11  float64 `json:"Minor_Tier1_1 "`
				MinorTier12  float64 `json:"Minor_Tier1_2 "`
				MinorTier13  float64 `json:"Minor_Tier1_3 "`
				MinorTier14  float64 `json:"Minor_Tier1_4 "`
				MinorTier15  float64 `json:"Minor_Tier1_5 "`
				MinorTier16  float64 `json:"Minor_Tier1_6 "`
				MinorTier17  float64 `json:"Minor_Tier1_7 "`
				MinorTier18  float64 `json:"Minor_Tier1_8 "`
				MinorTier19  float64 `json:"Minor_Tier1_9 "`
				MinorTier110 float64 `json:"Minor_Tier1_10 "`
				MinorTier111 float64 `json:"Minor_Tier1_11 "`
				MinorTier112 float64 `json:"Minor_Tier1_12 "`
				MinorTier113 float64 `json:"Minor_Tier1_13 "`
				MinorTier114 float64 `json:"Minor_Tier1_14 "`
				MinorTier115 float64 `json:"Minor_Tier1_15 "`
				MinorTier116 float64 `json:"Minor_Tier1_16 "`
				MinorTier117 float64 `json:"Minor_Tier1_17 "`
				MinorTier118 float64 `json:"Minor_Tier1_18 "`
				MinorTier119 float64 `json:"Minor_Tier1_19 "`
				MinorTier120 float64 `json:"Minor_Tier1_20 "`
				MinorTier21  float64 `json:"Minor_Tier2_1 "`
				MinorTier22  float64 `json:"Minor_Tier2_2 "`
				MinorTier23  float64 `json:"Minor_Tier2_3 "`
				MinorTier24  float64 `json:"Minor_Tier2_4 "`
				MinorTier25  float64 `json:"Minor_Tier2_5 "`
				MinorTier26  float64 `json:"Minor_Tier2_6 "`
				MinorTier27  float64 `json:"Minor_Tier2_7 "`
				MinorTier28  float64 `json:"Minor_Tier2_8 "`
				MinorTier29  float64 `json:"Minor_Tier2_9 "`
				MinorTier210 float64 `json:"Minor_Tier2_10 "`
				MinorTier211 float64 `json:"Minor_Tier2_11 "`
				MinorTier212 float64 `json:"Minor_Tier2_12 "`
				MinorTier213 float64 `json:"Minor_Tier2_13 "`
				MinorTier214 float64 `json:"Minor_Tier2_14 "`
				MinorTier215 float64 `json:"Minor_Tier2_15 "`
				MinorTier216 float64 `json:"Minor_Tier2_16 "`
				MinorTier217 float64 `json:"Minor_Tier2_17 "`
				MinorTier218 float64 `json:"Minor_Tier2_18 "`
				MinorTier219 float64 `json:"Minor_Tier2_19 "`
				MinorTier220 float64 `json:"Minor_Tier2_20 "`
				MinorTier221 float64 `json:"Minor_Tier2_21 "`
				MinorTier222 float64 `json:"Minor_Tier2_22 "`
				MinorTier223 float64 `json:"Minor_Tier2_23 "`
				MinorTier224 float64 `json:"Minor_Tier2_24 "`
				MinorTier225 float64 `json:"Minor_Tier2_25 "`
				MinorTier226 float64 `json:"Minor_Tier2_26 "`
				MinorTier227 float64 `json:"Minor_Tier2_27 "`
				MinorTier228 float64 `json:"Minor_Tier2_28 "`
				MinorTier229 float64 `json:"Minor_Tier2_29 "`
				MinorTier230 float64 `json:"Minor_Tier2_30 "`
			} `json:"Core"`
			Uncommon struct {
				Major1       float64 `json:"Major_1 "`
				Major2       float64 `json:"Major_2 "`
				Major3       float64 `json:"Major_3 "`
				Major4       float64 `json:"Major_4 "`
				MinorTier11  float64 `json:"Minor_Tier1_1 "`
				MinorTier12  float64 `json:"Minor_Tier1_2 "`
				MinorTier13  float64 `json:"Minor_Tier1_3 "`
				MinorTier14  float64 `json:"Minor_Tier1_4 "`
				MinorTier15  float64 `json:"Minor_Tier1_5 "`
				MinorTier16  float64 `json:"Minor_Tier1_6 "`
				MinorTier17  float64 `json:"Minor_Tier1_7 "`
				MinorTier18  float64 `json:"Minor_Tier1_8 "`
				MinorTier19  float64 `json:"Minor_Tier1_9 "`
				MinorTier110 float64 `json:"Minor_Tier1_10 "`
				MinorTier111 float64 `json:"Minor_Tier1_11 "`
				MinorTier112 float64 `json:"Minor_Tier1_12 "`
				MinorTier113 float64 `json:"Minor_Tier1_13 "`
				MinorTier114 float64 `json:"Minor_Tier1_14 "`
				MinorTier115 float64 `json:"Minor_Tier1_15 "`
				MinorTier116 float64 `json:"Minor_Tier1_16 "`
				MinorTier117 float64 `json:"Minor_Tier1_17 "`
				MinorTier118 float64 `json:"Minor_Tier1_18 "`
				MinorTier119 float64 `json:"Minor_Tier1_19 "`
				MinorTier120 float64 `json:"Minor_Tier1_20 "`
				MinorTier21  float64 `json:"Minor_Tier2_1 "`
				MinorTier22  float64 `json:"Minor_Tier2_2 "`
				MinorTier23  float64 `json:"Minor_Tier2_3 "`
				MinorTier24  float64 `json:"Minor_Tier2_4 "`
				MinorTier25  float64 `json:"Minor_Tier2_5 "`
				MinorTier26  float64 `json:"Minor_Tier2_6 "`
				MinorTier27  float64 `json:"Minor_Tier2_7 "`
				MinorTier28  float64 `json:"Minor_Tier2_8 "`
				MinorTier29  float64 `json:"Minor_Tier2_9 "`
				MinorTier210 float64 `json:"Minor_Tier2_10 "`
				MinorTier211 float64 `json:"Minor_Tier2_11 "`
				MinorTier212 float64 `json:"Minor_Tier2_12 "`
				MinorTier213 float64 `json:"Minor_Tier2_13 "`
				MinorTier214 float64 `json:"Minor_Tier2_14 "`
				MinorTier215 float64 `json:"Minor_Tier2_15 "`
				MinorTier216 float64 `json:"Minor_Tier2_16 "`
				MinorTier217 float64 `json:"Minor_Tier2_17 "`
				MinorTier218 float64 `json:"Minor_Tier2_18 "`
				MinorTier219 float64 `json:"Minor_Tier2_19 "`
				MinorTier220 float64 `json:"Minor_Tier2_20 "`
				MinorTier221 float64 `json:"Minor_Tier2_21 "`
				MinorTier222 float64 `json:"Minor_Tier2_22 "`
				MinorTier223 float64 `json:"Minor_Tier2_23 "`
				MinorTier224 float64 `json:"Minor_Tier2_24 "`
				MinorTier225 float64 `json:"Minor_Tier2_25 "`
				MinorTier226 float64 `json:"Minor_Tier2_26 "`
				MinorTier227 float64 `json:"Minor_Tier2_27 "`
				MinorTier228 float64 `json:"Minor_Tier2_28 "`
				MinorTier229 float64 `json:"Minor_Tier2_29 "`
				MinorTier230 float64 `json:"Minor_Tier2_30 "`
			} `json:"Uncommon"`
			Rare struct {
				Major1       float64 `json:"Major_1 "`
				Major2       float64 `json:"Major_2 "`
				Major3       float64 `json:"Major_3 "`
				Major4       float64 `json:"Major_4 "`
				MinorTier11  float64 `json:"Minor_Tier1_1 "`
				MinorTier12  float64 `json:"Minor_Tier1_2 "`
				MinorTier13  float64 `json:"Minor_Tier1_3 "`
				MinorTier14  float64 `json:"Minor_Tier1_4 "`
				MinorTier15  float64 `json:"Minor_Tier1_5 "`
				MinorTier16  float64 `json:"Minor_Tier1_6 "`
				MinorTier17  float64 `json:"Minor_Tier1_7 "`
				MinorTier18  float64 `json:"Minor_Tier1_8 "`
				MinorTier19  float64 `json:"Minor_Tier1_9 "`
				MinorTier110 float64 `json:"Minor_Tier1_10 "`
				MinorTier111 float64 `json:"Minor_Tier1_11 "`
				MinorTier112 float64 `json:"Minor_Tier1_12 "`
				MinorTier113 float64 `json:"Minor_Tier1_13 "`
				MinorTier114 float64 `json:"Minor_Tier1_14 "`
				MinorTier115 float64 `json:"Minor_Tier1_15 "`
				MinorTier116 float64 `json:"Minor_Tier1_16 "`
				MinorTier117 float64 `json:"Minor_Tier1_17 "`
				MinorTier118 float64 `json:"Minor_Tier1_18 "`
				MinorTier119 float64 `json:"Minor_Tier1_19 "`
				MinorTier120 float64 `json:"Minor_Tier1_20 "`
				MinorTier21  float64 `json:"Minor_Tier2_1 "`
				MinorTier22  float64 `json:"Minor_Tier2_2 "`
				MinorTier23  float64 `json:"Minor_Tier2_3 "`
				MinorTier24  float64 `json:"Minor_Tier2_4 "`
				MinorTier25  float64 `json:"Minor_Tier2_5 "`
				MinorTier26  float64 `json:"Minor_Tier2_6 "`
				MinorTier27  float64 `json:"Minor_Tier2_7 "`
				MinorTier28  float64 `json:"Minor_Tier2_8 "`
				MinorTier29  float64 `json:"Minor_Tier2_9 "`
				MinorTier210 float64 `json:"Minor_Tier2_10 "`
				MinorTier211 float64 `json:"Minor_Tier2_11 "`
				MinorTier212 float64 `json:"Minor_Tier2_12 "`
				MinorTier213 float64 `json:"Minor_Tier2_13 "`
				MinorTier214 float64 `json:"Minor_Tier2_14 "`
				MinorTier215 float64 `json:"Minor_Tier2_15 "`
				MinorTier216 float64 `json:"Minor_Tier2_16 "`
				MinorTier217 float64 `json:"Minor_Tier2_17 "`
				MinorTier218 float64 `json:"Minor_Tier2_18 "`
				MinorTier219 float64 `json:"Minor_Tier2_19 "`
				MinorTier220 float64 `json:"Minor_Tier2_20 "`
				MinorTier221 float64 `json:"Minor_Tier2_21 "`
				MinorTier222 float64 `json:"Minor_Tier2_22 "`
				MinorTier223 float64 `json:"Minor_Tier2_23 "`
				MinorTier224 float64 `json:"Minor_Tier2_24 "`
				MinorTier225 float64 `json:"Minor_Tier2_25 "`
				MinorTier226 float64 `json:"Minor_Tier2_26 "`
				MinorTier227 float64 `json:"Minor_Tier2_27 "`
				MinorTier228 float64 `json:"Minor_Tier2_28 "`
				MinorTier229 float64 `json:"Minor_Tier2_29 "`
				MinorTier230 float64 `json:"Minor_Tier2_30 "`
			} `json:"Rare"`
		} `json:"standard pack"`
		PremiumPack struct {
			Core struct {
				Major1       float64 `json:"Major_1 "`
				Major2       float64 `json:"Major_2 "`
				Major3       float64 `json:"Major_3 "`
				Major4       float64 `json:"Major_4 "`
				MinorTier11  float64 `json:"Minor_Tier1_1 "`
				MinorTier12  float64 `json:"Minor_Tier1_2 "`
				MinorTier13  float64 `json:"Minor_Tier1_3 "`
				MinorTier14  float64 `json:"Minor_Tier1_4 "`
				MinorTier15  float64 `json:"Minor_Tier1_5 "`
				MinorTier16  float64 `json:"Minor_Tier1_6 "`
				MinorTier17  float64 `json:"Minor_Tier1_7 "`
				MinorTier18  float64 `json:"Minor_Tier1_8 "`
				MinorTier19  float64 `json:"Minor_Tier1_9 "`
				MinorTier110 float64 `json:"Minor_Tier1_10 "`
				MinorTier111 float64 `json:"Minor_Tier1_11 "`
				MinorTier112 float64 `json:"Minor_Tier1_12 "`
				MinorTier113 float64 `json:"Minor_Tier1_13 "`
				MinorTier114 float64 `json:"Minor_Tier1_14 "`
				MinorTier115 float64 `json:"Minor_Tier1_15 "`
				MinorTier116 float64 `json:"Minor_Tier1_16 "`
				MinorTier117 float64 `json:"Minor_Tier1_17 "`
				MinorTier118 float64 `json:"Minor_Tier1_18 "`
				MinorTier119 float64 `json:"Minor_Tier1_19 "`
				MinorTier120 float64 `json:"Minor_Tier1_20 "`
				MinorTier21  float64 `json:"Minor_Tier2_1 "`
				MinorTier22  float64 `json:"Minor_Tier2_2 "`
				MinorTier23  float64 `json:"Minor_Tier2_3 "`
				MinorTier24  float64 `json:"Minor_Tier2_4 "`
				MinorTier25  float64 `json:"Minor_Tier2_5 "`
				MinorTier26  float64 `json:"Minor_Tier2_6 "`
				MinorTier27  float64 `json:"Minor_Tier2_7 "`
				MinorTier28  float64 `json:"Minor_Tier2_8 "`
				MinorTier29  float64 `json:"Minor_Tier2_9 "`
				MinorTier210 float64 `json:"Minor_Tier2_10 "`
				MinorTier211 float64 `json:"Minor_Tier2_11 "`
				MinorTier212 float64 `json:"Minor_Tier2_12 "`
				MinorTier213 float64 `json:"Minor_Tier2_13 "`
				MinorTier214 float64 `json:"Minor_Tier2_14 "`
				MinorTier215 float64 `json:"Minor_Tier2_15 "`
				MinorTier216 float64 `json:"Minor_Tier2_16 "`
				MinorTier217 float64 `json:"Minor_Tier2_17 "`
				MinorTier218 float64 `json:"Minor_Tier2_18 "`
				MinorTier219 float64 `json:"Minor_Tier2_19 "`
				MinorTier220 float64 `json:"Minor_Tier2_20 "`
				MinorTier221 float64 `json:"Minor_Tier2_21 "`
				MinorTier222 float64 `json:"Minor_Tier2_22 "`
				MinorTier223 float64 `json:"Minor_Tier2_23 "`
				MinorTier224 float64 `json:"Minor_Tier2_24 "`
				MinorTier225 float64 `json:"Minor_Tier2_25 "`
				MinorTier226 float64 `json:"Minor_Tier2_26 "`
				MinorTier227 float64 `json:"Minor_Tier2_27 "`
				MinorTier228 float64 `json:"Minor_Tier2_28 "`
				MinorTier229 float64 `json:"Minor_Tier2_29 "`
				MinorTier230 float64 `json:"Minor_Tier2_30 "`
			} `json:"Core"`
			Uncommon struct {
				Major1       float64 `json:"Major_1 "`
				Major2       float64 `json:"Major_2 "`
				Major3       float64 `json:"Major_3 "`
				Major4       float64 `json:"Major_4 "`
				MinorTier11  float64 `json:"Minor_Tier1_1 "`
				MinorTier12  float64 `json:"Minor_Tier1_2 "`
				MinorTier13  float64 `json:"Minor_Tier1_3 "`
				MinorTier14  float64 `json:"Minor_Tier1_4 "`
				MinorTier15  float64 `json:"Minor_Tier1_5 "`
				MinorTier16  float64 `json:"Minor_Tier1_6 "`
				MinorTier17  float64 `json:"Minor_Tier1_7 "`
				MinorTier18  float64 `json:"Minor_Tier1_8 "`
				MinorTier19  float64 `json:"Minor_Tier1_9 "`
				MinorTier110 float64 `json:"Minor_Tier1_10 "`
				MinorTier111 float64 `json:"Minor_Tier1_11 "`
				MinorTier112 float64 `json:"Minor_Tier1_12 "`
				MinorTier113 float64 `json:"Minor_Tier1_13 "`
				MinorTier114 float64 `json:"Minor_Tier1_14 "`
				MinorTier115 float64 `json:"Minor_Tier1_15 "`
				MinorTier116 float64 `json:"Minor_Tier1_16 "`
				MinorTier117 float64 `json:"Minor_Tier1_17 "`
				MinorTier118 float64 `json:"Minor_Tier1_18 "`
				MinorTier119 float64 `json:"Minor_Tier1_19 "`
				MinorTier120 float64 `json:"Minor_Tier1_20 "`
				MinorTier21  float64 `json:"Minor_Tier2_1 "`
				MinorTier22  float64 `json:"Minor_Tier2_2 "`
				MinorTier23  float64 `json:"Minor_Tier2_3 "`
				MinorTier24  float64 `json:"Minor_Tier2_4 "`
				MinorTier25  float64 `json:"Minor_Tier2_5 "`
				MinorTier26  float64 `json:"Minor_Tier2_6 "`
				MinorTier27  float64 `json:"Minor_Tier2_7 "`
				MinorTier28  float64 `json:"Minor_Tier2_8 "`
				MinorTier29  float64 `json:"Minor_Tier2_9 "`
				MinorTier210 float64 `json:"Minor_Tier2_10 "`
				MinorTier211 float64 `json:"Minor_Tier2_11 "`
				MinorTier212 float64 `json:"Minor_Tier2_12 "`
				MinorTier213 float64 `json:"Minor_Tier2_13 "`
				MinorTier214 float64 `json:"Minor_Tier2_14 "`
				MinorTier215 float64 `json:"Minor_Tier2_15 "`
				MinorTier216 float64 `json:"Minor_Tier2_16 "`
				MinorTier217 float64 `json:"Minor_Tier2_17 "`
				MinorTier218 float64 `json:"Minor_Tier2_18 "`
				MinorTier219 float64 `json:"Minor_Tier2_19 "`
				MinorTier220 float64 `json:"Minor_Tier2_20 "`
				MinorTier221 float64 `json:"Minor_Tier2_21 "`
				MinorTier222 float64 `json:"Minor_Tier2_22 "`
				MinorTier223 float64 `json:"Minor_Tier2_23 "`
				MinorTier224 float64 `json:"Minor_Tier2_24 "`
				MinorTier225 float64 `json:"Minor_Tier2_25 "`
				MinorTier226 float64 `json:"Minor_Tier2_26 "`
				MinorTier227 float64 `json:"Minor_Tier2_27 "`
				MinorTier228 float64 `json:"Minor_Tier2_28 "`
				MinorTier229 float64 `json:"Minor_Tier2_29 "`
				MinorTier230 float64 `json:"Minor_Tier2_30 "`
			} `json:"Uncommon"`
			Rare struct {
				Major1       float64 `json:"Major_1 "`
				Major2       float64 `json:"Major_2 "`
				Major3       float64 `json:"Major_3 "`
				Major4       float64 `json:"Major_4 "`
				MinorTier11  float64 `json:"Minor_Tier1_1 "`
				MinorTier12  float64 `json:"Minor_Tier1_2 "`
				MinorTier13  float64 `json:"Minor_Tier1_3 "`
				MinorTier14  float64 `json:"Minor_Tier1_4 "`
				MinorTier15  float64 `json:"Minor_Tier1_5 "`
				MinorTier16  float64 `json:"Minor_Tier1_6 "`
				MinorTier17  float64 `json:"Minor_Tier1_7 "`
				MinorTier18  float64 `json:"Minor_Tier1_8 "`
				MinorTier19  float64 `json:"Minor_Tier1_9 "`
				MinorTier110 float64 `json:"Minor_Tier1_10 "`
				MinorTier111 float64 `json:"Minor_Tier1_11 "`
				MinorTier112 float64 `json:"Minor_Tier1_12 "`
				MinorTier113 float64 `json:"Minor_Tier1_13 "`
				MinorTier114 float64 `json:"Minor_Tier1_14 "`
				MinorTier115 float64 `json:"Minor_Tier1_15 "`
				MinorTier116 float64 `json:"Minor_Tier1_16 "`
				MinorTier117 float64 `json:"Minor_Tier1_17 "`
				MinorTier118 float64 `json:"Minor_Tier1_18 "`
				MinorTier119 float64 `json:"Minor_Tier1_19 "`
				MinorTier120 float64 `json:"Minor_Tier1_20 "`
				MinorTier21  float64 `json:"Minor_Tier2_1 "`
				MinorTier22  float64 `json:"Minor_Tier2_2 "`
				MinorTier23  float64 `json:"Minor_Tier2_3 "`
				MinorTier24  float64 `json:"Minor_Tier2_4 "`
				MinorTier25  float64 `json:"Minor_Tier2_5 "`
				MinorTier26  float64 `json:"Minor_Tier2_6 "`
				MinorTier27  float64 `json:"Minor_Tier2_7 "`
				MinorTier28  float64 `json:"Minor_Tier2_8 "`
				MinorTier29  float64 `json:"Minor_Tier2_9 "`
				MinorTier210 float64 `json:"Minor_Tier2_10 "`
				MinorTier211 float64 `json:"Minor_Tier2_11 "`
				MinorTier212 float64 `json:"Minor_Tier2_12 "`
				MinorTier213 float64 `json:"Minor_Tier2_13 "`
				MinorTier214 float64 `json:"Minor_Tier2_14 "`
				MinorTier215 float64 `json:"Minor_Tier2_15 "`
				MinorTier216 float64 `json:"Minor_Tier2_16 "`
				MinorTier217 float64 `json:"Minor_Tier2_17 "`
				MinorTier218 float64 `json:"Minor_Tier2_18 "`
				MinorTier219 float64 `json:"Minor_Tier2_19 "`
				MinorTier220 float64 `json:"Minor_Tier2_20 "`
				MinorTier221 float64 `json:"Minor_Tier2_21 "`
				MinorTier222 float64 `json:"Minor_Tier2_22 "`
				MinorTier223 float64 `json:"Minor_Tier2_23 "`
				MinorTier224 float64 `json:"Minor_Tier2_24 "`
				MinorTier225 float64 `json:"Minor_Tier2_25 "`
				MinorTier226 float64 `json:"Minor_Tier2_26 "`
				MinorTier227 float64 `json:"Minor_Tier2_27 "`
				MinorTier228 float64 `json:"Minor_Tier2_28 "`
				MinorTier229 float64 `json:"Minor_Tier2_29 "`
				MinorTier230 float64 `json:"Minor_Tier2_30 "`
			} `json:"Rare"`
		} `json:"premium pack"`
		ElitePack struct {
			Core struct {
				Major1       float64 `json:"Major_1 "`
				Major2       float64 `json:"Major_2 "`
				Major3       float64 `json:"Major_3 "`
				Major4       float64 `json:"Major_4 "`
				MinorTier11  float64 `json:"Minor_Tier1_1 "`
				MinorTier12  float64 `json:"Minor_Tier1_2 "`
				MinorTier13  float64 `json:"Minor_Tier1_3 "`
				MinorTier14  float64 `json:"Minor_Tier1_4 "`
				MinorTier15  float64 `json:"Minor_Tier1_5 "`
				MinorTier16  float64 `json:"Minor_Tier1_6 "`
				MinorTier17  float64 `json:"Minor_Tier1_7 "`
				MinorTier18  float64 `json:"Minor_Tier1_8 "`
				MinorTier19  float64 `json:"Minor_Tier1_9 "`
				MinorTier110 float64 `json:"Minor_Tier1_10 "`
				MinorTier111 float64 `json:"Minor_Tier1_11 "`
				MinorTier112 float64 `json:"Minor_Tier1_12 "`
				MinorTier113 float64 `json:"Minor_Tier1_13 "`
				MinorTier114 float64 `json:"Minor_Tier1_14 "`
				MinorTier115 float64 `json:"Minor_Tier1_15 "`
				MinorTier116 float64 `json:"Minor_Tier1_16 "`
				MinorTier117 float64 `json:"Minor_Tier1_17 "`
				MinorTier118 float64 `json:"Minor_Tier1_18 "`
				MinorTier119 float64 `json:"Minor_Tier1_19 "`
				MinorTier120 float64 `json:"Minor_Tier1_20 "`
				MinorTier21  float64 `json:"Minor_Tier2_1 "`
				MinorTier22  float64 `json:"Minor_Tier2_2 "`
				MinorTier23  float64 `json:"Minor_Tier2_3 "`
				MinorTier24  float64 `json:"Minor_Tier2_4 "`
				MinorTier25  float64 `json:"Minor_Tier2_5 "`
				MinorTier26  float64 `json:"Minor_Tier2_6 "`
				MinorTier27  float64 `json:"Minor_Tier2_7 "`
				MinorTier28  float64 `json:"Minor_Tier2_8 "`
				MinorTier29  float64 `json:"Minor_Tier2_9 "`
				MinorTier210 float64 `json:"Minor_Tier2_10 "`
				MinorTier211 float64 `json:"Minor_Tier2_11 "`
				MinorTier212 float64 `json:"Minor_Tier2_12 "`
				MinorTier213 float64 `json:"Minor_Tier2_13 "`
				MinorTier214 float64 `json:"Minor_Tier2_14 "`
				MinorTier215 float64 `json:"Minor_Tier2_15 "`
				MinorTier216 float64 `json:"Minor_Tier2_16 "`
				MinorTier217 float64 `json:"Minor_Tier2_17 "`
				MinorTier218 float64 `json:"Minor_Tier2_18 "`
				MinorTier219 float64 `json:"Minor_Tier2_19 "`
				MinorTier220 float64 `json:"Minor_Tier2_20 "`
				MinorTier221 float64 `json:"Minor_Tier2_21 "`
				MinorTier222 float64 `json:"Minor_Tier2_22 "`
				MinorTier223 float64 `json:"Minor_Tier2_23 "`
				MinorTier224 float64 `json:"Minor_Tier2_24 "`
				MinorTier225 float64 `json:"Minor_Tier2_25 "`
				MinorTier226 float64 `json:"Minor_Tier2_26 "`
				MinorTier227 float64 `json:"Minor_Tier2_27 "`
				MinorTier228 float64 `json:"Minor_Tier2_28 "`
				MinorTier229 float64 `json:"Minor_Tier2_29 "`
				MinorTier230 float64 `json:"Minor_Tier2_30 "`
			} `json:"Core"`
			Uncommon struct {
				Major1       float64 `json:"Major_1 "`
				Major2       float64 `json:"Major_2 "`
				Major3       float64 `json:"Major_3 "`
				Major4       float64 `json:"Major_4 "`
				MinorTier11  float64 `json:"Minor_Tier1_1 "`
				MinorTier12  float64 `json:"Minor_Tier1_2 "`
				MinorTier13  float64 `json:"Minor_Tier1_3 "`
				MinorTier14  float64 `json:"Minor_Tier1_4 "`
				MinorTier15  float64 `json:"Minor_Tier1_5 "`
				MinorTier16  float64 `json:"Minor_Tier1_6 "`
				MinorTier17  float64 `json:"Minor_Tier1_7 "`
				MinorTier18  float64 `json:"Minor_Tier1_8 "`
				MinorTier19  float64 `json:"Minor_Tier1_9 "`
				MinorTier110 float64 `json:"Minor_Tier1_10 "`
				MinorTier111 float64 `json:"Minor_Tier1_11 "`
				MinorTier112 float64 `json:"Minor_Tier1_12 "`
				MinorTier113 float64 `json:"Minor_Tier1_13 "`
				MinorTier114 float64 `json:"Minor_Tier1_14 "`
				MinorTier115 float64 `json:"Minor_Tier1_15 "`
				MinorTier116 float64 `json:"Minor_Tier1_16 "`
				MinorTier117 float64 `json:"Minor_Tier1_17 "`
				MinorTier118 float64 `json:"Minor_Tier1_18 "`
				MinorTier119 float64 `json:"Minor_Tier1_19 "`
				MinorTier120 float64 `json:"Minor_Tier1_20 "`
				MinorTier21  float64 `json:"Minor_Tier2_1 "`
				MinorTier22  float64 `json:"Minor_Tier2_2 "`
				MinorTier23  float64 `json:"Minor_Tier2_3 "`
				MinorTier24  float64 `json:"Minor_Tier2_4 "`
				MinorTier25  float64 `json:"Minor_Tier2_5 "`
				MinorTier26  float64 `json:"Minor_Tier2_6 "`
				MinorTier27  float64 `json:"Minor_Tier2_7 "`
				MinorTier28  float64 `json:"Minor_Tier2_8 "`
				MinorTier29  float64 `json:"Minor_Tier2_9 "`
				MinorTier210 float64 `json:"Minor_Tier2_10 "`
				MinorTier211 float64 `json:"Minor_Tier2_11 "`
				MinorTier212 float64 `json:"Minor_Tier2_12 "`
				MinorTier213 float64 `json:"Minor_Tier2_13 "`
				MinorTier214 float64 `json:"Minor_Tier2_14 "`
				MinorTier215 float64 `json:"Minor_Tier2_15 "`
				MinorTier216 float64 `json:"Minor_Tier2_16 "`
				MinorTier217 float64 `json:"Minor_Tier2_17 "`
				MinorTier218 float64 `json:"Minor_Tier2_18 "`
				MinorTier219 float64 `json:"Minor_Tier2_19 "`
				MinorTier220 float64 `json:"Minor_Tier2_20 "`
				MinorTier221 float64 `json:"Minor_Tier2_21 "`
				MinorTier222 float64 `json:"Minor_Tier2_22 "`
				MinorTier223 float64 `json:"Minor_Tier2_23 "`
				MinorTier224 float64 `json:"Minor_Tier2_24 "`
				MinorTier225 float64 `json:"Minor_Tier2_25 "`
				MinorTier226 float64 `json:"Minor_Tier2_26 "`
				MinorTier227 float64 `json:"Minor_Tier2_27 "`
				MinorTier228 float64 `json:"Minor_Tier2_28 "`
				MinorTier229 float64 `json:"Minor_Tier2_29 "`
				MinorTier230 float64 `json:"Minor_Tier2_30 "`
			} `json:"Uncommon"`
			Rare struct {
				Major1       float64 `json:"Major_1 "`
				Major2       float64 `json:"Major_2 "`
				Major3       float64 `json:"Major_3 "`
				Major4       float64 `json:"Major_4 "`
				MinorTier11  float64 `json:"Minor_Tier1_1 "`
				MinorTier12  float64 `json:"Minor_Tier1_2 "`
				MinorTier13  float64 `json:"Minor_Tier1_3 "`
				MinorTier14  float64 `json:"Minor_Tier1_4 "`
				MinorTier15  float64 `json:"Minor_Tier1_5 "`
				MinorTier16  float64 `json:"Minor_Tier1_6 "`
				MinorTier17  float64 `json:"Minor_Tier1_7 "`
				MinorTier18  float64 `json:"Minor_Tier1_8 "`
				MinorTier19  float64 `json:"Minor_Tier1_9 "`
				MinorTier110 float64 `json:"Minor_Tier1_10 "`
				MinorTier111 float64 `json:"Minor_Tier1_11 "`
				MinorTier112 float64 `json:"Minor_Tier1_12 "`
				MinorTier113 float64 `json:"Minor_Tier1_13 "`
				MinorTier114 float64 `json:"Minor_Tier1_14 "`
				MinorTier115 float64 `json:"Minor_Tier1_15 "`
				MinorTier116 float64 `json:"Minor_Tier1_16 "`
				MinorTier117 float64 `json:"Minor_Tier1_17 "`
				MinorTier118 float64 `json:"Minor_Tier1_18 "`
				MinorTier119 float64 `json:"Minor_Tier1_19 "`
				MinorTier120 float64 `json:"Minor_Tier1_20 "`
				MinorTier21  float64 `json:"Minor_Tier2_1 "`
				MinorTier22  float64 `json:"Minor_Tier2_2 "`
				MinorTier23  float64 `json:"Minor_Tier2_3 "`
				MinorTier24  float64 `json:"Minor_Tier2_4 "`
				MinorTier25  float64 `json:"Minor_Tier2_5 "`
				MinorTier26  float64 `json:"Minor_Tier2_6 "`
				MinorTier27  float64 `json:"Minor_Tier2_7 "`
				MinorTier28  float64 `json:"Minor_Tier2_8 "`
				MinorTier29  float64 `json:"Minor_Tier2_9 "`
				MinorTier210 float64 `json:"Minor_Tier2_10 "`
				MinorTier211 float64 `json:"Minor_Tier2_11 "`
				MinorTier212 float64 `json:"Minor_Tier2_12 "`
				MinorTier213 float64 `json:"Minor_Tier2_13 "`
				MinorTier214 float64 `json:"Minor_Tier2_14 "`
				MinorTier215 float64 `json:"Minor_Tier2_15 "`
				MinorTier216 float64 `json:"Minor_Tier2_16 "`
				MinorTier217 float64 `json:"Minor_Tier2_17 "`
				MinorTier218 float64 `json:"Minor_Tier2_18 "`
				MinorTier219 float64 `json:"Minor_Tier2_19 "`
				MinorTier220 float64 `json:"Minor_Tier2_20 "`
				MinorTier221 float64 `json:"Minor_Tier2_21 "`
				MinorTier222 float64 `json:"Minor_Tier2_22 "`
				MinorTier223 float64 `json:"Minor_Tier2_23 "`
				MinorTier224 float64 `json:"Minor_Tier2_24 "`
				MinorTier225 float64 `json:"Minor_Tier2_25 "`
				MinorTier226 float64 `json:"Minor_Tier2_26 "`
				MinorTier227 float64 `json:"Minor_Tier2_27 "`
				MinorTier228 float64 `json:"Minor_Tier2_28 "`
				MinorTier229 float64 `json:"Minor_Tier2_29 "`
				MinorTier230 float64 `json:"Minor_Tier2_30 "`
			} `json:"Rare"`
		} `json:"elite pack"`
	} `json:"Trigger detail"`
}
