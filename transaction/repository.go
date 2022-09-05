package transaction

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetCampaignByID(campaignID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaignByID(campaignID int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transaction).Error
	return transaction, err
}

func (r *repository) GetByUserID(userID int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	return transactions, err
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	return transaction, err
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	return transaction, err
}
