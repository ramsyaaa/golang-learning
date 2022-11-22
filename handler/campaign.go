package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) CampaignInput(c *gin.Context) {
	var input campaign.CampaignInput

	err := c.ShouldBindJSON(&input)

	if(err != nil) { 
		
		errors := helper.ErrorValidationFormat(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Create new Campaign failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCampaign, err := h.campaignService.CampaignInput(input)

	if(err != nil) { 
		response := helper.APIResponse("Create new Campaign failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatUser(newCampaign)

	response := helper.APIResponse("Campaign is Successfully Created", http.StatusOK, "Success", formatter)
	c.JSON(http.StatusOK, response)
}