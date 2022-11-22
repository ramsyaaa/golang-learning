package campaign

type CampaignFormatter struct {
	UserID int `json:"id"`
	Name string `json:"name"`
	ShortDescription string `json:"ShortDescription"`
	Description string `json:"description"`
	Perks string `json:"perks"`
	BackerCount int `json:"backer_count"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	Slug string `json:"slug"`

}

func FormatUser(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		UserID: campaign.UserID,
		Name:   campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		Perks:         campaign.Perks,
		BackerCount:    campaign.BackerCount,
		GoalAmount:    campaign.GoalAmount,
		CurrentAmount:  campaign.CurrentAmount,
		Slug:          campaign.Slug,
	}

	return formatter
}