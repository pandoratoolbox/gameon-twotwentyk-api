package venly

const (
	VENLY_API_URL    = "https://api-wallet-sandbox.venly.io"
	VENLY_AUTH_URL   = "https://login-sandbox.venly.io"
	VENLY_CLIENT_ID  = "3e572e5a-8c4d-466f-a674-02a030f9aa57"
	VENLY_APP_SECRET = "2971ea69-564e-4c71-833d-f29d9217495b"
)

type VenlyRequestCreateWallet struct {
	Description string `json:"description,omitempty"`
	PinCode     string `json:"pinCode,omitempty"`
	SecretType  string `json:"secretType,omitempty"`
	WalletType  string `json:"walletType,omitem"`
}

func Login() (string, error) {
	var token string
	return token, nil
}

// func CreateWallet() {
// 	//create wallet
// 	//store wallet in db
// 	//return wallet
// 	url := VENLY_API_URL + "/api/wallets"
// 	res, err := http.Post(url, "application/json", nil)
// }
