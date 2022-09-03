package handler

import "time"

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     int
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
