package setup

import (
	controllerSmartContract "backend/controller/smartcontract"
	"github.com/gin-gonic/gin"
)

func SetupEndpoints(router *gin.Engine) {
	smartContractController := router.Group("/smartcontract")
	smartContractController.GET("/balance/:walletAddress", controllerSmartContract.GetBalanceHandler())
	smartContractController.POST("/addition/:smartContractAddress", controllerSmartContract.PostAdditionFunctionHandler())
}
