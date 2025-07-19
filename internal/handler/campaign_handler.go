package handler

import "berbagi/internal/services"

type campaignHandler struct {
	campaignService services.Service
}

func NewCampaignHandler(campaignService services.Service) *campaignHandler {
	return &campaignHandler{campaignService, }
}