package app

import (
	"Task1/internal/controllers/rest"
	"Task1/internal/domain"
	"Task1/internal/repository"
	"Task1/internal/services"
	"os"
)

type DependenciesManager struct {
	linkPairRepository domain.LinkPairRepository
	linkPairService    domain.LinkPairService
	linkPairHandler    domain.LinkPairHandler
}

func NewDependenciesManager(useDB bool) *DependenciesManager {
	var linkPairRepository domain.LinkPairRepository

	if !useDB {
		linkPairRepository = repository.NewLinkPairRepository()
	} else {

		dbHost := os.Getenv("DATABASE_HOST")
		dbPort := os.Getenv("DATABASE_PORT")
		dbUser := os.Getenv("DATABASE_USER")
		dbPass := os.Getenv("DATABASE_PASS")
		dbName := os.Getenv("DATABASE_NAME")
		createQuary := os.Getenv("CREATE_TABLE")

		linkPairRepository = repository.NewLinkPairRepositoryDB(dbHost, dbPort, dbUser, dbPass, dbName, createQuary)
	}

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	linkPairService := services.NewLinkPairService(linkPairRepository, serverHost, serverPort)
	return &DependenciesManager{
		linkPairRepository: linkPairRepository,
		linkPairService:    linkPairService,
		linkPairHandler:    rest.NewLinkPairHandler(linkPairService),
	}

}
