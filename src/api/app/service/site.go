package service

import (
	"website-conector/app/domain"
)

type SiteService struct {
	repository domain.SiteRepository
}

func NewSiteService(repository domain.SiteRepository) *SiteService {
	return &SiteService{
		repository: repository,
	}
}

func (s *SiteService) GetSites() ([]domain.SiteDTO, error) {
	return s.repository.GetSites()
}

func (s *SiteService) GetSitesWithLogin() ([]domain.SiteLoginDTO, error) {
	return s.repository.GetSitesWithLogin()
}
