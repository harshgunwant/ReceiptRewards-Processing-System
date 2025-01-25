package handlers

import (
	"FetchRewardsAssessment/internal/services"
	"FetchRewardsAssessment/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPoints(c *gin.Context) {
	id := c.Param("id")

	points, err := services.GetPoints(id)
	if err != nil {
		utils.SendError(c, http.StatusNotFound, "No receipt found for that ID.", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
