package handlers

import (
	"eth-wallet-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateWallet godoc
// @Summary Create a new Ethereum wallet
// @Description Create a new Ethereum wallet and return the address and private key
// @Tags ethereum
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /eth/create-wallet [post]
func CreateWallet(c *gin.Context) {
	address, privateKey, mnemonic, err := services.CreateWallet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"address": address, "privateKey": privateKey, "mnemonic": mnemonic})
}

// MultiSigTransfer godoc
// @Summary Perform a multi-signature transfer
// @Description Perform a multi-signature transfer of Ethereum
// @Tags ethereum
// @Accept  json
// @Produce  json
// @Param   body body services.MultiSigTransferRequest true "Multi-signature transfer details"
// @Success 200 {object} map[string]string
// @Router /eth/multi-sig-transfer [post]
func MultiSigTransfer(c *gin.Context) {
	var req services.MultiSigTransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := services.MultiSigTransfer(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"txHash": txHash})
}

// GetContractBalance godoc
// @Summary Get the balance of a contract
// @Description Get the balance of a specified contract address
// @Tags ethereum
// @Accept  json
// @Produce  json
// @Param   address query string true "Contract address"
// @Success 200 {object} map[string]string
// @Router /eth/contract-balance [get]
func GetContractBalance(c *gin.Context) {
	address := c.Query("address")
	balance, err := services.GetContractBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

// GetGasBalance godoc
// @Summary Get the gas balance of an address
// @Description Get the gas (ETH) balance of a specified address
// @Tags ethereum
// @Accept  json
// @Produce  json
// @Param   address query string true "Ethereum address"
// @Success 200 {object} map[string]string
// @Router /eth/gas-balance [get]
func GetGasBalance(c *gin.Context) {
	address := c.Query("address")
	balance, err := services.GetGasBalance(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}

// BurnNFT godoc
// @Summary Burn an NFT
// @Description Burn (destroy) a specified NFT
// @Tags ethereum
// @Accept  json
// @Produce  json
// @Param   body body services.BurnNFTRequest true "NFT burn details"
// @Success 200 {object} map[string]string
// @Router /eth/burn-nft [post]
func BurnNFT(c *gin.Context) {
	var req services.BurnNFTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := services.BurnNFT(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"txHash": txHash})
}

// GetTokenSupply godoc
// @Summary Get the total supply of a token
// @Description Get the total supply of a specified ERC20 token
// @Tags ethereum
// @Accept  json
// @Produce  json
// @Param   address query string true "Token contract address"
// @Success 200 {object} map[string]string
// @Router /eth/token-supply [get]
func GetTokenSupply(c *gin.Context) {
	address := c.Query("address")
	supply, err := services.GetTokenSupply(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"supply": supply})
}

// Withdraw godoc
// @Summary Withdraw tokens
// @Description Withdraw tokens from the service to a specified address
// @Tags ethereum
// @Accept  json
// @Produce  json
// @Param   body body services.WithdrawRequest true "Withdrawal details"
// @Success 200 {object} map[string]string
// @Router /eth/withdraw [post]
func Withdraw(c *gin.Context) {
	var req services.WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txHash, err := services.Withdraw(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"txHash": txHash})
}
