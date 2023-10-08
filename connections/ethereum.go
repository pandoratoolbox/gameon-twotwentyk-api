package connections

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	cardpackABI "gameon-twotwentyk-api/bindings/CardPack"
	collectionABI "gameon-twotwentyk-api/bindings/Collection"
	craftingABI "gameon-twotwentyk-api/bindings/Crafting"
	factoriesABI "gameon-twotwentyk-api/bindings/Factories"
	openPackABI "gameon-twotwentyk-api/bindings/OpenPack"
	partsABI "gameon-twotwentyk-api/bindings/Parts"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionResult struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	Hash    string `json:"result"`
}

type CollectionMetaType struct {
	Name   string      `json:"name"`
	Config interface{} `json:"config"`
	Src    string      `json:"src"`
}

type PackMetaType struct {
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config interface{} `json:config`
	Src    string      `json:"src"`
}

type CardMetaType struct {
	Name   string `json:"name"`
	Tier   string `json:"tier"`
	Rarity string `json:"rarity"`
	Src    string `json:"src"`
}

type CollectionType struct {
	Metadata CollectionMetaType
	Id       int64
	TokenId  *big.Int
	Address  common.Address
	Packs    *[]PackType
}

type PackType struct {
	Metadata PackMetaType
	Id       int64
	TokenId  *big.Int
	Address  common.Address
	Cards    *[]CardType
}

type CardType struct {
	Metadata     CardMetaType
	Id           int64
	TokenId      *big.Int
	Type         string
	Address      common.Address
	CollectionId int64
	PackId       int64
}

type PackConfig []int64
type Packs []PackConfig

var (
	// rpc       = "https://polygon-mumbai.g.alchemy.com/v2/PTi4_upQduxYasRED-Hy1EBldW7pBstt"
	rpc       = os.Getenv("ETHEREUM_URL")
	walletKey = os.Getenv("WALLET_PRIVATE_KEY")

	value    = big.NewInt(0)
	gasLimit = uint64(20_000_000)
	chainId  = big.NewInt(0)

	ownerAddr = common.HexToAddress("0xf22679BBaf587B9c776D0A25a05e64B214f19CAd")

	collectionAddr  = common.HexToAddress(os.Getenv("COLLECTION_ADDRESS"))
	packFactoryAddr = common.HexToAddress(os.Getenv("CARDPACKFACTORY_ADDRESS"))

	categoryFactoryAddr = common.HexToAddress(os.Getenv("CATEGORYFACTORY_ADDRESS"))
	yearFactoryAddr     = common.HexToAddress(os.Getenv("YEARFACTORY_ADDRESS"))

	dayMonthFactoryAddr = common.HexToAddress(os.Getenv("DAYMONTHFACTORY_ADDRESS"))
	triggerFactoryAddr  = common.HexToAddress(os.Getenv("TRIGGERFACTORY_ADDRESS"))

	craftingCardFactoryAddr = common.HexToAddress(os.Getenv("CRAFTINGCARDFACTORY_ADDRESS"))

	// identityFactoryAddr     = common.HexToAddress(os.Getenv("YEARFACTORY_ADDRESS"))
	// predictionFactoryAddr   = common.HexToAddress(os.Getenv("YEARFACTORY_ADDRESS"))
	packAddress = common.HexToAddress(os.Getenv("CARDPACKFACTORY_ADDRESS"))

	categoryAddress        common.Address
	yearAddress            common.Address
	dayMonthAddress        common.Address
	craftingCardAddress    common.Address
	triggerAddress         common.Address
	identityAddress        common.Address
	predictionAddress      common.Address
	collectionCreatedTopic = "0x0ba92c0f760258e3285b70eb8220efa72578d8bb68dfa3e26d48f3cecd2e10b7"

	collectionId = big.NewInt(0)
	totalPackAmt = big.NewInt(300 + 200 + 100)

	quantity         = big.NewInt(0)
	standardPackTier = big.NewInt(0)

	packOpenerAddr = common.HexToAddress("0x5D6A8216590cEEeFd149f3CB26ad096B8516bB00")

	collectionInstance          *collectionABI.Collection
	cardPackFactoryInstance     *factoriesABI.CardPackFactory
	categoryFactoryInstance     *factoriesABI.CategoryFactory
	yearFactoryInstance         *factoriesABI.YearFactory
	dayMonthFactoryInstance     *factoriesABI.DayMonthFactory
	craftingCardFactoryInstance *factoriesABI.CraftingCardFactory
	triggerFactoryInstance      *factoriesABI.TriggerFactory
	// identityFactoryInstance     *factoriesABI.IdentityFactory
	// predictionFactoryInstance   *factoriesABI.PredictionFactory

	cardpackInstance     *cardpackABI.CardPack
	categoryInstance     *partsABI.Category
	yearInstance         *partsABI.Year
	daymonthInstance     *partsABI.DayMonth
	craftingcardInstance *partsABI.CraftingCard
	triggerInstance      *partsABI.Trigger
	PackOpenerInstance   *openPackABI.PackOpener
	IdentityInstance     *craftingABI.CollectionCrafting
	predictionInstance   *craftingABI.PredictionCrafting

	cards      []CardType
	pack       PackType
	collection CollectionType

	err error
)

func InitEthereum() {
	_chainID, err := strconv.ParseInt(os.Getenv("CHAIN_ID"), 10, 64)
	fmt.Printf("chainID:", _chainID)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	chainId = big.NewInt(_chainID)

	// setCollectionMinterAndAdmin()
	// CreateCollection()
	// getAddresses(collectionId)
	// setMinterAndCrafter()
	// createCard()
	// openPack()
	// craftingIdentity()
	// craftingPrediction()
}

func setCollectionMinterAndAdmin() {
	client, _ := ethclient.Dial(rpc)
	var (
		auth *bind.TransactOpts
	)

	collectionInstance, err = collectionABI.NewCollection(collectionAddr, client)
	auth = getTransactor()
	collectionInstance.ChangeMinter(auth, ownerAddr)

	cardPackFactoryInstance, err = factoriesABI.NewCardPackFactory(packFactoryAddr, client)
	auth = getTransactor()
	cardPackFactoryInstance.ChangeAdmin(auth, collectionAddr)

	categoryFactoryInstance, err = factoriesABI.NewCategoryFactory(categoryFactoryAddr, client)
	auth = getTransactor()
	categoryFactoryInstance.ChangeAdmin(auth, collectionAddr)

	yearFactoryInstance, err = factoriesABI.NewYearFactory(yearFactoryAddr, client)
	auth = getTransactor()
	yearFactoryInstance.ChangeAdmin(auth, collectionAddr)

	dayMonthFactoryInstance, err = factoriesABI.NewDayMonthFactory(dayMonthFactoryAddr, client)
	auth = getTransactor()
	dayMonthFactoryInstance.ChangeAdmin(auth, collectionAddr)

	craftingCardFactoryInstance, err = factoriesABI.NewCraftingCardFactory(craftingCardFactoryAddr, client)
	auth = getTransactor()
	craftingCardFactoryInstance.ChangeAdmin(auth, collectionAddr)

	triggerFactoryInstance, err = factoriesABI.NewTriggerFactory(triggerFactoryAddr, client)
	auth = getTransactor()
	triggerFactoryInstance.ChangeAdmin(auth, collectionAddr)

	// identityFactoryInstance, err = factoriesABI.NewIdentityFactory(identityFactoryAddr, client)
	// auth = getTransactor()
	// identityFactoryInstance.ChangeAdmin(auth, collectionAddr)

	// predictionFactoryInstance, err = factoriesABI.NewPredictionFactory(predictionFactoryAddr, client)
	// auth = getTransactor()
	// predictionFactoryInstance.ChangeAdmin(auth, collectionAddr)
}
func CreateCollection() {

	client, _ := ethclient.Dial(rpc)
	auth := getTransactor()

	fmt.Println(totalPackAmt)
	collectionInstance, err = collectionABI.NewCollection(collectionAddr, client)
	tx, _ := collectionInstance.CreateCollection(auth,
		"https://example.com/collection/0", totalPackAmt)

	fmt.Printf("create collection tx hash: %s", tx.Hash())

}

func getAddresses(collectionId *big.Int) (common.Address, common.Address, common.Address, common.Address, common.Address, common.Address, common.Address, common.Address, error) {
	client, _ := ethclient.Dial(rpc)

	var (
		logs []types.Log
	)

	collectionABI, err := collectionABI.CollectionMetaData.GetAbi()

	query := ethereum.FilterQuery{
		Addresses: []common.Address{collectionAddr},
		Topics:    [][]common.Hash{{common.HexToHash(collectionCreatedTopic)}},
	}

	logs, err = client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
		return common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, err
	}

	for _, vlog := range logs {
		topic := vlog.Topics[0].Hex()
		fmt.Println(topic)
		if topic == collectionCreatedTopic {

			data, _ := collectionABI.Unpack("CollectionCreated", vlog.Data)
			// spew.Dump("topic", data)

			tokenId := data[0].(*big.Int)
			fmt.Println(tokenId, collectionId, tokenId.Cmp(collectionId))
			if tokenId.Cmp(collectionId) == 0 {
				packAddress = data[1].(common.Address)
				categoryAddress = data[2].(common.Address)
				yearAddress = data[3].(common.Address)
				dayMonthAddress = data[4].(common.Address)
				craftingCardAddress = data[5].(common.Address)
				triggerAddress = data[6].(common.Address)
				identityAddress = data[7].(common.Address)
				predictionAddress = data[8].(common.Address)

				fmt.Println("card pack address", packAddress)
				fmt.Println("category address", categoryAddress)
				fmt.Println("year address", yearAddress)
				fmt.Println("daymonth address", dayMonthAddress)
				fmt.Println("trigger address", triggerAddress)
				fmt.Println("craftingcard address", craftingCardAddress)
				fmt.Println("identity address", identityAddress)
				fmt.Println("prediction address", predictionAddress)

				return packAddress, categoryAddress, yearAddress, dayMonthAddress, craftingCardAddress, triggerAddress, identityAddress, predictionAddress, err
			}
		}
	}
	return common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, common.Address{}, err
}

func setMinterAndCrafter() {
	client, _ := ethclient.Dial(rpc)
	var (
		auth *bind.TransactOpts
	)

	cardpackInstance, err = cardpackABI.NewCardPack(packAddress, client)
	auth = getTransactor()
	cardpackInstance.ChangeMinter(auth, ownerAddr)

	auth = getTransactor()
	cardpackInstance.ChangeOpener(auth, packOpenerAddr)

	categoryInstance, err = partsABI.NewCategory(categoryAddress, client)
	auth = getTransactor()
	categoryInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	categoryInstance.ChangeCrafter(auth, identityAddress)

	yearInstance, err = partsABI.NewYear(yearAddress, client)
	auth = getTransactor()
	yearInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	yearInstance.ChangeCrafter(auth, identityAddress)

	daymonthInstance, err = partsABI.NewDayMonth(dayMonthAddress, client)
	auth = getTransactor()
	daymonthInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	daymonthInstance.ChangeCrafter(auth, identityAddress)

	craftingcardInstance, err = partsABI.NewCraftingCard(craftingCardAddress, client)
	auth = getTransactor()
	craftingcardInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	craftingcardInstance.ChangeCrafter(auth, identityAddress)
	auth = getTransactor()
	craftingcardInstance.ChangePredictor(auth, predictionAddress)

	triggerInstance, err = partsABI.NewTrigger(triggerAddress, client)
	auth = getTransactor()
	triggerInstance.ChangeMinter(auth, packOpenerAddr)
	auth = getTransactor()
	triggerInstance.ChangeCrafter(auth, predictionAddress)

	// IdentifyInstance, err = factoriesABI.NewIdentity(identityAddress, client)
	// auth = getTransactor()
	// IdentifyInstance.ChangeCrafter(auth, predictionAddress)

}

func createCard() {
	var (
		tokenURI = "https://example.com/collecition/0/cardpack/0"
	)
	client, _ := ethclient.Dial(rpc)
	cardpackInstance, err = cardpackABI.NewCardPack(packAddress, client)
	auth := getTransactor()
	tx, _ := cardpackInstance.Mint(auth, ownerAddr, quantity, standardPackTier, tokenURI)
	fmt.Printf("card pack tx hash: %s", tx.Hash())
}

func openPack() {
	var (
		cardPackTokenId = big.NewInt(0)

		categoryTokenURIs     = []string{"https://example.com/collection/0/category/1"}
		yearTOkenURIs         = []string{"https://example.com/collection/0/year/1"}
		daymonthTokenURIs     = []string{"https://example.com/collection/0/daymonth/1"}
		craftingCardTokenURIs = []string{
			"https://example.com/collection/0/craftingcard/1",
			"https://example.com/collection/0/craftingcard/2",
			"https://example.com/collection/0/craftingcard/3",
		}
		triggerTokenURIs = []string{"https://example.com/collection/0/trigger/1"}
	)

	client, _ := ethclient.Dial(rpc)
	PackOpenerInstance, err = openPackABI.NewPackOpener(packOpenerAddr, client)
	auth := getTransactor()

	tx, _ := PackOpenerInstance.OpenPack(
		auth,
		cardPackTokenId,
		packAddress,
		categoryAddress,
		yearAddress,
		dayMonthAddress,
		craftingCardAddress,
		triggerAddress,
		categoryTokenURIs,
		yearTOkenURIs,
		daymonthTokenURIs,
		craftingCardTokenURIs,
		triggerTokenURIs,
	)

	fmt.Printf("open pack tx hash: %s", tx.Hash())
}

func craftingIdentity() {
	client, _ := ethclient.Dial(rpc)
	var (
		categoryTokenId = big.NewInt(0)
		dayMonthTokenId = big.NewInt(0)
		yearTokenId     = big.NewInt(0)
		craftingTokenId = big.NewInt(1)

		identityTokenUrl = "https://example.com/collection/0/identity/0"
	)

	IdentityInstance, err = craftingABI.NewCollectionCrafting(predictionAddress, client)
	auth := getTransactor()
	tx, _ := IdentityInstance.CraftCollection(auth,
		dayMonthAddress, yearAddress, categoryAddress, craftingCardAddress,
		dayMonthTokenId, yearTokenId, categoryTokenId, craftingTokenId, identityTokenUrl)

	fmt.Printf("crafting identify tx hash: %s", tx.Hash())
}

func craftingPrediction() {
	client, _ := ethclient.Dial(rpc)
	var (
		identityTokenId    = big.NewInt(0)
		triggerTokenIds    = []*big.Int{big.NewInt(0)}
		craftingTokenId    = big.NewInt(2)
		predictionTokenUrl = "https://example.com/collection/0/prediction/0"
	)

	predictionInstance, err = craftingABI.NewPredictionCrafting(predictionAddress, client)
	auth := getTransactor()
	tx, _ := predictionInstance.CraftCollection(auth,
		identityAddress, triggerAddress, craftingCardAddress,
		identityTokenId, triggerTokenIds, craftingTokenId, predictionTokenUrl)

	fmt.Printf("tx sent: %s", tx.Hash())
}

func getTransactor() *bind.TransactOpts {
	client, _ := ethclient.Dial(rpc)

	privatekey, _ := crypto.HexToECDSA(walletKey)

	gasPrice, err := client.SuggestGasPrice(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// gasPrice, _ := client.SuggestGasPrice(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privatekey, chainId)
	if err == nil {
		fmt.Print("Auth not created")
	}

	// auth := bind.NewKeyedTransactor(privatekey)
	auth.Nonce = big.NewInt(int64(getNonce(privatekey)))
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	signfunc := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.NewEIP155Signer(chainId)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privatekey)
		if err != nil {
			fmt.Printf("Error signing: %v\n", err)
			os.Exit(1)
			return nil, nil
		}
		return tx.WithSignature(signer, signature)
	}
	auth.Signer = signfunc

	return auth
}

func getNonce(privatekey *ecdsa.PrivateKey) uint64 {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privatekey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, _ := client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Println("from", fromAddress, nonce)

	return nonce
}
