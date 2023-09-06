package venly

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	VENLY_API_URL    = "https://api-wallet-sandbox.venly.io"
	VENLY_AUTH_URL   = "https://login-sandbox.venly.io"
	VENLY_CLIENT_ID  = "3e572e5a-8c4d-466f-a674-02a030f9aa57"
	VENLY_APP_SECRET = "2971ea69-564e-4c71-833d-f29d9217495b"
)

var Global *VenlyClient

type VenlyRequestCreateWallet struct {
	Description string `json:"description,omitempty"`
	PinCode     string `json:"pinCode,omitempty"`
	SecretType  string `json:"secretType,omitempty"`
	WalletType  string `json:"walletType,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
}

type VenlyWallet struct {
	ID           string `json:"id"`
	Address      string `json:"address"`
	WalletType   string `json:"walletType"`
	SecretType   string `json:"secretType"`
	CreatedAt    string `json:"createdAt"`
	Archived     bool   `json:"archived"`
	Description  string `json:"description"`
	Primary      bool   `json:"primary"`
	HasCustomPin bool   `json:"hasCustomPin"`
	Identifier   string `json:"identifier"`
	Balance      struct {
		Available     bool   `json:"available"`
		SecretType    string `json:"secretType"`
		Balance       int    `json:"balance"`
		GasBalance    int    `json:"gasBalance"`
		Symbol        string `json:"symbol"`
		GasSymbol     string `json:"gasSymbol"`
		RawBalance    string `json:"rawBalance"`
		RawGasBalance string `json:"rawGasBalance"`
		Decimals      int    `json:"decimals"`
	} `json:"balance"`
}

type VenlyResponseCreateWallet struct {
	Success bool        `json:"success"`
	Result  VenlyWallet `json:"result"`
}

type VenlyRequestGetAccessToken struct {
	ClientId     string
	ClientSecret string
	GrantType    string
}

type VenlyResponseGetAccessToken struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type VenlyClient struct {
	Http        http.Client
	Config      VenlyClientConfig
	AccessToken string
	ExpiresAt   time.Time
}

type VenlyClientConfig struct {
	ClientId     string
	ClientSecret string
}

func NewClient(config VenlyClientConfig) (*VenlyClient, error) {
	client := VenlyClient{
		Config: config,
		Http:   *http.DefaultClient,
	}

	return &client, nil
}

func (c *VenlyClient) SendRequest(t string, url string, data interface{}) ([]byte, error) {
	if time.Now().After(c.ExpiresAt) {
		err := c.GetAccessToken()
		if err != nil {
			return nil, err
		}
	}

	if t == "POST" {
		if data != nil {
			js, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}

			r, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
			if err != nil {
				return nil, err
			}

			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}

			return b, nil
		}

		r, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return nil, err
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		return b, nil
	}

	if t == "GET" {
		r, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}

		return b, nil

	}

	return nil, errors.New("Invalid request type")
}

func (c *VenlyClient) GetAccessToken() error {
	url := VENLY_AUTH_URL + "/auth/realms/Arkane/protocol/openid-connect/token"
	data := VenlyRequestGetAccessToken{
		ClientId:     c.Config.ClientId,
		ClientSecret: c.Config.ClientSecret,
		GrantType:    "client_credentials",
	}

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")

	var resp VenlyResponseGetAccessToken
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&resp)
	if err != nil {
		return err
	}

	c.Http.Transport = &http.Transport{}

	c.AccessToken = resp.AccessToken

	return nil
}

func (c *VenlyClient) CreateWallet(data VenlyRequestCreateWallet) (VenlyWallet, error) {
	var wallet VenlyWallet
	url := VENLY_API_URL + "/api/wallets"

	// data := VenlyRequestCreateWallet{
	// 	Description: "test",
	// 	PinCode:     "1234",
	// 	SecretType:  "MATIC",
	// 	WalletType:  "WHITE_LABEL",
	// 	Identifier:  "type=unrecoverable",
	// }

	var resp VenlyResponseCreateWallet

	b, err := c.SendRequest("POST", url, data)
	if err != nil {
		return wallet, err
	}

	err = json.Unmarshal(b, &resp)
	if err != nil {
		return wallet, err
	}

	if !resp.Success {
		return wallet, errors.New("failed to create wallet")
	}

	wallet = resp.Result

	return wallet, nil
}

type VenlyRequestTransferNft struct {
	Pincode            int `json:"pincode"`
	TransactionRequest struct {
		Type         string `json:"type"`
		SecretType   string `json:"secretType"`
		WalletID     string `json:"walletId"`
		To           string `json:"to"`
		TokenAddress string `json:"tokenAddress"`
		TokenID      int    `json:"tokenId"`
	} `json:"transactionRequest"`
}

type VenlyResponseTransferNft struct {
}

// func (c *VenlyClient) TransferNft(data VenlyRequestTransferNft) (VenlyResponseTransferNft, error) {
// 	var res VenlyResponseTransferNft
// 	url := VENLY_API_URL + "/api/transactions/execute"

// 	// {
// 	// 	"pincode": 1234,
// 	// 	"transactionRequest": {
// 	// 		"type": "NFT_TRANSFER",
// 	// 		"secretType": "MATIC",
// 	// 		"walletId": "89733529-ddc1-44aa-b2af-17601e2e26c0",
// 	// 		"to": "0x9d9376EEbFE3443544d3654f7aD272f0A31D8152",
// 	// 		"tokenAddress": "0xd05a795d339886bb8dd46cfe2ac009d7f1e48a64",
// 	// 		"tokenId": 91
// 	// 	}
// 	// }
// 	return res, nil
// }
