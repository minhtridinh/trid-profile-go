package repository

import (
	"github.com/dinhminhtri/triD-profile/internal/model"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) GetByID(id uint) (*model.Profile, error) {
	var profile model.Profile
	if err := r.db.First(&profile, id).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) Create(profile *model.Profile) error {
	return r.db.Create(profile).Error
}

func (r *ProfileRepository) Update(profile *model.Profile) error {
	return r.db.Save(profile).Error
}
