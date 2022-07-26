package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcellribeiro/tunity_test/pkg"
	"net/http"
	"strconv"
)

func GetCollatz(c *gin.Context) {
	number := c.Param("number")
	floatNumber, err := strconv.ParseFloat(number, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := pkg.GetCollatz(floatNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
