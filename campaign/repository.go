package campaign

import "gorm.io/gorm"

type Repository interface {
	Save(campaign Campaign) (Campaign, error)
}

type repository struct { 
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error

	if(err != nil) {
		return campaign, err
	}

	return campaign, nil
}