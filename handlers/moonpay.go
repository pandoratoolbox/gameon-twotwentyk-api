package handlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gameon-twotwentyk-api/models"
	"gameon-twotwentyk-api/store"
	"net/http"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
)

const (
	MOONPAY_SECRET_KEY      = "sk_test_cHMS4IN3BLApimnz8C0RYtyObpbOT4H"
	MOONPAY_WEBHOOK_KEY     = "wk_test_BWUpCi6aQXMPoGo5yKA8ZkjRjlFvD60J"
	MOONPAY_PUBLISHABLE_KEY = "pk_test_PaUTi3HVAHvclaZTMJS0TNTfMIrpPj"
)

type MoonpayNft struct {
	TokenId           string  `json:"tokenId,omitempty"`
	ContractAddress   string  `json:"contractAddress,omitempty"`
	Name              string  `json:"name,omitempty"`
	Collection        string  `json:"collection,omitempty"`
	ImageUrl          string  `json:"imageUrl,omitempty"`
	ExplorerUrl       string  `json:"explorerUrl,omitempty"`
	Price             float64 `json:"price,omitempty"`
	PriceCurrencyCode string  `json:"priceCurrencyCode,omitempty"`
	Quantity          int64   `json:"quantity,omitempty"`
	SellerAddress     string  `json:"sellerAddress,omitempty"`
	SellType          string  `json:"sellType,omitempty"`
	Flow              string  `json:"flow,omitempty"`
	Network           string  `json:"network,omitempty"`
}

func WebhookMoonpayGetNftInfo(w http.ResponseWriter, r *http.Request) {
	var out MoonpayNft
	ctx := r.Context()

	contract_address := chi.URLParam(r, "contract_address")

	token_id_q := chi.URLParam(r, "token_id")

	//0 doesn't pass through json encoding - fix in custom encoderS
	token_id := int64(14)

	t_secondary := MoonpayNft{
		TokenId:           token_id_q,
		ContractAddress:   contract_address,
		PriceCurrencyCode: "USDC",
		SellType:          "Secondary",
		Flow:              "Lite",
		Network:           "polygon",
		Quantity:          1,
		// ImageUrl:          "https://www.debate.com.mx/__export/1491334879891/sites/debate/img/2017/04/04/macaco_crop1491334754748.jpg_423682103.jpg",
	}

	t_primary := MoonpayNft{
		TokenId:           "__1", //return same tokenId but save transaction in our database with real new generated tokenID
		ContractAddress:   contract_address,
		PriceCurrencyCode: "USDC",
		SellType:          "Primary",
		Flow:              "Lite",
		Network:           "polygon",
		Quantity:          1,
		Price:             1,
		Collection:        "genesis",
		Name:              "elite card pack",
		// ImageUrl:          "https://www.debate.com.mx/__export/1491334879891/sites/debate/img/2017/04/04/macaco_crop1491334754748.jpg_423682103.jpg",
	}

	if !strings.Contains(token_id_q, "_") {
		out = t_secondary
		var err error
		token_id, err = strconv.ParseInt(token_id_q, 10, 64)
		if err != nil {
			ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

	} else {
		out = t_primary
	}

	// listing_id_q := r.URL.Query().Get("listingId")
	// buyer_wallet_address := r.URL.Query().Get("walletAddress")

	// nft_collection, ok := store.GetNftCollectionByContractAddress(ctx, contract_address)
	// if !ok {
	// 	ServeError(w, fmt.Sprintf("Unable to find nft collection for contract address (%s)", contract_address), http.StatusBadRequest)
	// 	return
	// }

	nft_collection := struct {
		Id   int64
		Name string
	}{
		Id:   models.NFT_TYPE_ID_CARD_PACK,
		Name: "card packs",
	}

	switch nft_collection.Id {
	case models.NFT_TYPE_ID_CARD_PACK:
		break
	case models.NFT_TYPE_ID_CATEGORY:
		listings, err := store.SearchMarketplaceListings(ctx, "", 1, []int64{nft_collection.Id}, 1, token_id, true)
		if err != nil {
			ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		listing := listings[0]

		nft, err := store.GetNftCardCategory(ctx, token_id)
		if err != nil {
			ServeError(w, err.Error(), http.StatusBadRequest)
			return
		}

		rarity, ok := store.NftRarityMap[*nft.Rarity]
		if !ok {
			ServeError(w, fmt.Sprintf("Error getting rarity from map with id (%d)", rarity), http.StatusBadRequest)
			return
		}

		out.Name = fmt.Sprintf("%s %s Category Card", *nft.Category, rarity)
		out.Collection = nft_collection.Name
		out.SellerAddress = *listing.Owner.WalletAddress
		out.Price = float64(*listing.Price)
		break
	case models.NFT_TYPE_ID_CRAFTING:
	case models.NFT_TYPE_ID_DAY_MONTH:
	case models.NFT_TYPE_ID_IDENTITY:
	case models.NFT_TYPE_ID_PREDICTION:
	case models.NFT_TYPE_ID_TRIGGER:
	case models.NFT_TYPE_ID_YEAR:
	default:
		ServeError(w, "NFT collection for given id is not currently supported.", http.StatusServiceUnavailable)
		return
	}

	// ServeJSON(w, out)

	js, err := json.Marshal(out)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	// w.Write(js)

	fmt.Println(string(js))

	// w.Write([]byte(`{
	// 	"tokenId":"109",
	// 	"contractAddress":"0x2180c6ecf2f770bd51dbb0d779cc81048899",
	// 	"name":"MoonRocket",
	// 	"collection":"MoonPay Special Collection",
	// 	"imageUrl":"https://b1ic5m3wqxz9zd.cloudfront.net/f793.jpg",
	// 	"explorerUrl":"",
	// 	"price":0.1,
	// 	"priceCurrencyCode":"ETH",
	// 	"quantity":1,
	// 	"sellerAddress":"0xt246e19c76a23068fa235e1673c10opecfbeb7hf",
	// 	"sellType":"Primary",
	// 	"flow":"Lite",
	// 	"network":"Ethereum"
	//  }`))

	w.Write(js)

}

func WebhookMoonpayTransaction(w http.ResponseWriter, r *http.Request) {
	input := json.RawMessage{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	spew.Dump(input)

	w.WriteHeader(200)
}

// func WebhookMoonpayNotifyDeliverNft(w http.ResponseWriter, r *http.Request) {
// 	contract_address := chi.URLParam(r, "contract_address")
// 	token_id := chi.URLParam(r, "token_id")

// 	input := struct {
// 		Mode                string
// 		BuyerWalletAddress  string
// 		PriceCurrencyCode   string
// 		Price               int64
// 		Quantity            string
// 		SellerWalletAddress string
// 		ListingId           string
// 	}{}

// 	w.WriteHeader(200)

// }

// func WebhookMoonpayGetNftTransactionStatus(w http.ResponseWriter, r *http.Request) {
// 	ids_q := r.URL.Query().Get("id")
// 	ids := strings.Split(ids_q, ",")

// 	out := struct {
// 		Id              string
// 		Status          string
// 		TransactionHash string
// 		StatusChangedAt time.Time
// 		failureReason   string
// 		tokenId         string
// 	}{}
// }

func MoonpaySignUrl(w http.ResponseWriter, r *http.Request) {
	input := struct {
		Url string `json:"url"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		ServeError(w, err.Error(), 500)
		return
	}

	fmt.Println("URL: " + input.Url)

	// host := strings.Split(input.Url, "?")[0]
	query := "?" + strings.Split(input.Url, "?")[1]
	secret := MOONPAY_SECRET_KEY
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(query))

	// Get result and encode as hexadecimal string
	// sha := hex.EncodeToString()

	// sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	sha := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Println("Signature: " + sha)
	ServeJSON(w, struct {
		Signature string
	}{
		Signature: sha,
	})
}
