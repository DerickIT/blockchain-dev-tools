package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/openweb3/go-ethereum-hdwallet"
)

type MultiSigTransferRequest struct {
	From    string   `json:"from"`
	To      string   `json:"to"`
	Amount  *big.Int `json:"amount"`
	Signers []string `json:"signers"`
}

type BurnNFTRequest struct {
	ContractAddress string   `json:"contractAddress"`
	TokenID         *big.Int `json:"tokenId"`
}

type WithdrawRequest struct {
	To     string   `json:"to"`
	Amount *big.Int `json:"amount"`
}

const infuraURL = "https://mainnet.infura.io/v3/YOUR-PROJECT-ID" // 替换为你的 Infura 项目 ID

func CreateWallet() (address, privateKey, mnemonic string, err error) {
	// 生成新的助记词
	entropy, err := hdwallet.NewEntropy(256)
	if err != nil {
		return "", "", "", err
	}

	mnemonic, err = hdwallet.NewMnemonicFromEntropy(entropy)
	if err != nil {
		return "", "", "", err
	}

	// 从助记词创建钱包
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return "", "", "", err
	}

	// 生成 HD 路径。这里使用以太坊的标准路径
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")

	// 从路径派生账户
	account, err := wallet.Derive(path, false)
	if err != nil {
		return "", "", "", err
	}

	// 获取私钥
	privateKeyECDSA, err := wallet.PrivateKey(account)
	if err != nil {
		return "", "", "", err
	}
	privateKey = common.Bytes2Hex(privateKeyECDSA.D.Bytes())

	// 获取地址
	address = account.Address.Hex()

	return address, privateKey, mnemonic, nil
}

func MultiSigTransfer(req MultiSigTransferRequest) (txHash string, err error) {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return "", err
	}

	fromAddress := common.HexToAddress(req.From)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(req.To)
	tx := types.NewTransaction(nonce, toAddress, req.Amount, 21000, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	for _, signer := range req.Signers {
		privateKey, err := crypto.HexToECDSA(signer)
		if err != nil {
			return "", err
		}
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			return "", err
		}
		tx = signedTx
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func GetContractBalance(address string) (balance *big.Int, err error) {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return nil, err
	}

	account := common.HexToAddress(address)
	balance, err = client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func GetGasBalance(address string) (balance *big.Int, err error) {
	return GetContractBalance(address)
}

func BurnNFT(req BurnNFTRequest) (txHash string, err error) {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return "", err
	}

	contractABI, err := abi.JSON(strings.NewReader("YOUR_CONTRACT_ABI"))
	if err != nil {
		return "", err
	}

	contractAddress := common.HexToAddress(req.ContractAddress)
	contract := bind.NewBoundContract(contractAddress, contractABI, client, client, client)

	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY") // 替换为你的私钥
	if err != nil {
		return "", err
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return "", err
	}

	tx, err := contract.Transact(auth, "burn", req.TokenID)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func GetTokenSupply(address string) (supply *big.Int, err error) {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return nil, err
	}

	contractABI, err := abi.JSON(strings.NewReader("YOUR_ERC20_CONTRACT_ABI"))
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(address)
	contract := bind.NewBoundContract(contractAddress, contractABI, client, client, client)

	var result []interface{}
	err = contract.Call(&bind.CallOpts{}, &result, "totalSupply")
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return big.NewInt(0), nil
	}

	supply, ok := result[0].(*big.Int)
	if !ok {
		return nil, errors.New("failed to convert result to *big.Int")
	}

	return supply, nil
}

func Withdraw(req WithdrawRequest) (txHash string, err error) {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY") // 替换为你的私钥
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(req.To)
	tx := types.NewTransaction(nonce, toAddress, req.Amount, 21000, gasPrice, nil)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
