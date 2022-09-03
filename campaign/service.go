package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID == 0 {
		campaign, err := s.repository.FindAll()
		return campaign, err
	}

	campaign, err := s.repository.FindByUserID(userID)
	return campaign, err
}

func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.ID)

	return campaign, err
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID
	campaign.Slug = getSlug(fmt.Sprintf("%s %d", input.Name, input.User.ID))

	// tambahan by Me :D
	campaignExist, _ := s.repository.FindBySlug(campaign.Slug)
	if campaignExist.ID > 0 {
		return campaign, errors.New("Can not continue, this user has already made the same campaign name")
	}

	newCampaign, err := s.repository.Save(campaign)
	return newCampaign, err
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not authorized")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount
	// campaign.Slug = getSlug(fmt.Sprintf("%s %d", inputData.Name, inputData.User.ID))

	updatedCampaign, err := s.repository.Update(campaign)

	return updatedCampaign, err
}

func getSlug(slugCandidate string) string {
	return slug.Make(slugCandidate)
}
