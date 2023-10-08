package connections

// import (
// 	"crypto/ecdsa"
// 	"fmt"
// 	"log"
// 	"math/big"
// 	"os"
// 	"strconv"
// 	"strings"

// 	collection "gameon-twotwentyk-api/bindings/Collection"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/joho/godotenv"
// )

// var (
// 	client            *ethclient.Client
// 	collectionAddress common.Address
// 	cardpackAddress   common.Address
// 	identityAddress   common.Address
// 	predictionAddress common.Address
// 	accountKey        *ecdsa.PrivateKey
// )

// func InitEthereum() error {

// 	var err error
// 	if os.Getenv("ETHEREUM_URL") == "" {
// 		err := godotenv.Load()
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	eth_url := os.Getenv("ETHEREUM_URL")
// 	client, err = ethclient.Dial(eth_url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collectionAddress = common.HexToAddress(os.Getenv("COLLECTION_ADDRESS"))
// 	identityAddress = common.HexToAddress(os.Getenv("IDENTITY_ADDRESS"))
// 	cardpackAddress = common.HexToAddress(os.Getenv("CARDPACK_ADDRESS"))
// 	predictionAddress = common.HexToAddress(os.Getenv("PREDICTION_ADDRESS"))

// 	accountKey = privateKeyToECDSA(os.Getenv("ACCOUNT_PRIVATE_KEY"))
// 	return nil
// }

// func privateKeyToECDSA(privateKeyHex string) *ecdsa.PrivateKey {
// 	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

// 	// Parse the private key and sender address
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		log.Fatal("Failed to parse private key:", err)
// 	}
// 	return privateKey
// }

// func PrivateKeyToAddress(privateKeyHex string) common.Address {
// 	privateKey := privateKeyToECDSA(privateKeyHex)
// 	return crypto.PubkeyToAddress(privateKey.PublicKey)
// }

// func CreateTransactionOpt() *bind.TransactOpts {
// 	// Convert the string to an integer
// 	chainID, err := strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
// 	fmt.Printf("chainID:", chainID)
// 	if err != nil {
// 		fmt.Println("Error converting string to int:", err)
// 		return nil
// 	}

// 	_chainID := big.NewInt(chainID) // For Ethereum mainnet
// 	auth, err := bind.NewKeyedTransactorWithChainID(accountKey, _chainID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &bind.TransactOpts{
// 		From:     auth.From,
// 		Signer:   auth.Signer,
// 		GasPrice: big.NewInt(69722716), // example gas price
// 		GasLimit: uint64(3000000),      // example gas limit
// 	}

// }

// func CreateCollection(tokenURI string, amount int64) {

// 	instance, err := collection.NewCollection(collectionAddress, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_amount := new(big.Int)
// 	_amount.SetInt64(amount) // Set the value to 1234567890

// 	tx, err := instance.CreateCollection(CreateTransactionOpt(), tokenURI, _amount)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("TX: %v", tx.Hash())
// 	// instance.
// 	// // Call your contract's method
// 	// _, err := instance.FilterApprovalForAll(&bind.CallOpts{})
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// result.
// }
