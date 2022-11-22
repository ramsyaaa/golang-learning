package campaign

type Service interface {
	CampaignInput(input CampaignInput) (Campaign, error)
}

type service struct { 
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CampaignInput(input CampaignInput) (Campaign, error) { 
	campaign := Campaign{}
	campaign.UserID = input.UserID
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.BackerCount = input.BackerCount
	campaign.GoalAmount = input.GoalAmount
	campaign.CurrentAmount = input.CurrentAmount
	campaign.Slug = input.Slug

	newCampaign, err := s.repository.Save(campaign)

	if( err != nil ) { 
		return newCampaign, err
	}

	return newCampaign, nil

}