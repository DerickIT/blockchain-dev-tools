package main

import (
	_ "eth-wallet-service/docs"
	"eth-wallet-service/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Ethereum Wallet Service API
// @version 1.0
// @description This is a sample server for Ethereum wallet operations.
// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		eth := v1.Group("/eth")
		{
			eth.POST("/create-wallet", handlers.CreateWallet)
			eth.POST("/multi-sig-transfer", handlers.MultiSigTransfer)
			eth.GET("/contract-balance", handlers.GetContractBalance)
			eth.GET("/gas-balance", handlers.GetGasBalance)
			eth.POST("/burn-nft", handlers.BurnNFT)
			eth.GET("/token-supply", handlers.GetTokenSupply)
			eth.POST("/withdraw", handlers.Withdraw)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
