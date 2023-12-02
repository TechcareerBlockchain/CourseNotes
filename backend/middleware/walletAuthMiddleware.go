package middleware

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WalletAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		walletAddress := c.GetHeader("walletAddress")
		if common.IsHexAddress(walletAddress) {
			c.Next()
		} else {
			fmt.Println("Wallet is not correct")
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"status": false, "data": "Wallet is not correct"})
		}
	}
}
