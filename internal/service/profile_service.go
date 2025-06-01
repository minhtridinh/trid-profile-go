package service

import (
	"github.com/minhtridinh/trid-profile-go/internal/model"
	"github.com/minhtridinh/trid-profile-go/internal/repository"
)

type ProfileService struct {
	repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetProfile(id uint) (*model.Profile, error) {
	return s.repo.GetByID(id)
}

func (s *ProfileService) CreateProfile(profile *model.Profile) error {
	return s.repo.Create(profile)
}

func (s *ProfileService) UpdateProfile(profile *model.Profile) error {
	return s.repo.Update(profile)
}
