package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gameon-twotwentyk-api/graphql"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/chi"
)

const (
	MOONPAY_SECRET_KEY      = "sk_test_cHMS4IN3BLApimnz8C0RYtyObpbOT4H"
	MOONPAY_WEBHOOK_KEY     = "wk_test_BWUpCi6aQXMPoGo5yKA8ZkjRjlFvD60J"
	MOONPAY_PUBLISHABLE_KEY = "pk_test_PaUTi3HVAHvclaZTMJS0TNTfMIrpPj"
)

func WebhookMoonpayGetNftInfo(w http.ResponseWriter, r *http.Request) {

	contract_address := chi.URLParam(r, "contract_address")

	token_id_q := chi.URLParam(r, "token_id")

	body := getNFTMetadata(contract_address, token_id_q)

	w.Write(body)

}

func getNFTMetadata(contract_address string, token_id_q string) []byte {
	url := "https://deep-index.moralis.io/api/v2.2/nft/%s/%s?chain=polygon&format=decimal&normalizeMetadata=true&media_items=false"
	url = fmt.Sprintf(url, contract_address, token_id_q)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", "YOUR_API_KEY")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

/*
* Function to deliver NFT to a buyerAddress
* 200 Transfer Success
* 400 Delivery could not be initiated. Response may include a reason for our logs.
* 403 NFT not available for sale. MoonPay will cancel the customer transaction.
* 404 NFT not found. MoonPay will cancel the customer transaction.
* 500Something went wrong. MoonPay will cancel the customer transaction.
 */
func WebhookMoonpayDeliverNFT(w http.ResponseWriter, r *http.Request) {
	contract_address := chi.URLParam(r, "contract_address")

	token_id_q := chi.URLParam(r, "token_id")

	body := struct {
		Mode                string
		BuyerWalletAddress  string
		PriceCurrencyCode   string
		Price               float64
		Quantity            string
		SellerWalletAddress string
		ListingId           string
	}{}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(data, &body)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get user by wallet address
	user, err := graphql.GetUserByWalletAddress(r.Context(), body.SellerWalletAddress)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nftMetadataBytes := getNFTMetadata(contract_address, token_id_q)

	nftMetadata := struct {
		TokenAddress string `json:"token_address"`
		TokenID      string `json:"token_id"`
	}{}

	err = json.Unmarshal(nftMetadataBytes, &nftMetadata)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if nftMetadata.TokenAddress == "" {
		ServeError(w, fmt.Sprintf("NFT On Contract %s with ID %s Not Found", contract_address, token_id_q), http.StatusNotFound)
		return
	}

	// TODO: Check if for sale
	// TODO: if not for sale retrurn 403 error
	// TODO: Return transactionId after transfering the nft to seller
	url := "https://api-wallet.venly.io/api/transactions/execute"
	transferNftBodyStr := `{
  "transactionRequest" : {
    "type" : "NFT_TRANSFER",
    "walletId" : "%s", 
    "to" : "%s",
    "secretType" : "MATIC",
    "tokenAddress" : "%s",
    "tokenId" : %d
  }
  "pincode" : "1234"
}`
	token_id, err := strconv.ParseInt(nftMetadata.TokenID, 10, 64)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transferNftBodyStr = fmt.Sprintf(transferNftBodyStr, user.VenlyId, body.BuyerWalletAddress, nftMetadata.TokenAddress, token_id)
	transferNftBody := []byte(transferNftBodyStr)
	// walletId = user.venly_id
	// to = body.BuyerWalletAddress
	// tokenAddress = nftMetadata.TokenAddress

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(transferNftBody))
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resBodyStruct := struct {
		Success bool `json:"success"`
		Result  struct {
			TransactionHash string `json:"transactionHash"`
		} `json:"result"`
	}{}

	err = json.Unmarshal(resBody, &resBodyStruct)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	finalRes := `{
		transactionId:%s
	}`
	finalRes = fmt.Sprintf(finalRes, resBodyStruct.Result.TransactionHash)
	w.Write([]byte(finalRes))

}

func WebhookMoonpayTransactionStatus(w http.ResponseWriter, r *http.Request) {
	tx_ids := strings.Split(r.URL.Query().Get("id"), ",")
	url := "GET : https://api-wallet.venly.io/api/transactions/MATIC/%s/status"

	type localResponse struct {
		Id              string    `json:"id"`
		Status          string    `json:"status"`
		TransactionHash []string  `json:"transactionHash"`
		StatusChangedAt time.Time `json:"statusChangedAt"`
		FailureReason   string    `json:"failureReason"`
		Tokenid         []string  `json:"tokenid"`
	}

	response := make([]localResponse, 0, len(tx_ids))

	for _, id := range tx_ids {
		url = fmt.Sprint(url, id)
		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("Accept", "application/json")
		req.Header.Add("X-API-Key", "YOUR_API_KEY")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := localResponse{}
		err = json.Unmarshal(body, &resp)
		if err != nil {
			ServeError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response = append(response, resp)
	}

	data, err := json.Marshal(response)
	if err != nil {
		ServeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)

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
