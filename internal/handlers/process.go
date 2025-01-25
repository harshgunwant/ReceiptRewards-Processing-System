package handlers

import (
	"FetchRewardsAssessment/internal/models"
	"FetchRewardsAssessment/internal/services"
	"FetchRewardsAssessment/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	// Parse and validate JSON payload
	if err := c.ShouldBindJSON(&receipt); err != nil {
		utils.SendError(c, http.StatusBadRequest, "The receipt is invalid.", "Please verify input. Error: "+err.Error())
		return
	}

	id, err := services.ProcessReceipt(&receipt)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "The receipt is invalid.", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
