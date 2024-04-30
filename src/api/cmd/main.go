package main

import (
	"fmt"
	"os"
	"website-conector/app/service"
	"website-conector/infra/database/mssql"
	"website-conector/infra/logger"
	"website-conector/infra/webserver"
	"website-conector/utils"
)

func init() {
	utils.LoadVars()
	logger.InitLog()
}

func generateInjection() *service.SiteService {

	siteRepository := mssql.NewSiteRepository()
	siteService := service.NewSiteService(siteRepository)

	return siteService
}

func main() {
	siteService := generateInjection()
	router := webserver.NewHttpServer(os.Getenv("PORT"), siteService)

	fmt.Println("main", fmt.Sprintf("running server at: %s", os.Getenv("PORT")))
	router.SetupRouter()
}
