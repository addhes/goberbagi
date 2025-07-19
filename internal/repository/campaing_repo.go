package repository

import "gorm.io/gorm"

type CampaignRepository interface {
}

type campaignrepository struct {
	db *gorm.DB
}