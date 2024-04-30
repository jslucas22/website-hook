package mssql

import (
	"website-conector/app/domain"
)

type SiteRepository struct{}

func NewSiteRepository() *SiteRepository {
	return &SiteRepository{}
}

func (r *SiteRepository) GetSites() ([]domain.SiteDTO, error) {
	sites := []domain.SiteDTO{
		{URL: "https://site1.com"},
		{URL: "https://site2.com"},
		{URL: "https://site3.com"},
	}
	return sites, nil
}

func (r *SiteRepository) GetSitesWithLogin() ([]domain.SiteLoginDTO, error) {
	sites := []domain.SiteLoginDTO{
		{URL: "https://site1.com", Username: "user1", Password: "pass1"},
		{URL: "https://site2.com", Username: "user2", Password: "pass2"},
		{URL: "https://site3.com", Username: "user3", Password: "pass3"},
	}
	return sites, nil
}
