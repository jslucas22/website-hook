package domain

type SiteDTO struct {
	URL string `json:"url"`
}

type SiteLoginDTO struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SiteRepository interface {
	GetSites() ([]SiteDTO, error)
	GetSitesWithLogin() ([]SiteLoginDTO, error)
}
