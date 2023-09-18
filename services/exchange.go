package services

import (
	"exchange/modules"
	"exchange/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExchangeRate(c *gin.Context) {
	source := c.Query("source")
	target := c.Query("target")
	amount := c.Query("amount")

	if source == "" || target == "" || amount == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The parameters cannot be empty"})
		return
	}

	parseAmount := utils.RemoveSpecialChars(amount)
	if !utils.IsInteger(parseAmount) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The amount must be an integer"})
		return
	}

	currency := modules.Currency.Data
	if currency == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "The currency info is empty"})
		return
	}

	rate, ok := currency[source][target]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The currency is not support"})
		return
	}

	targetAmount, err := strconv.ParseFloat(parseAmount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	price := rate * targetAmount
	c.JSON(http.StatusOK, gin.H{
		"msg":    "success",
		"source": source,
		"target": target,
		"amount": utils.FormatFloatWithCommas(price),
	})
}
