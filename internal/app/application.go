package app

import (
	"github.com/gin-gonic/gin"
	"os"
)

type Application interface {
	Run() error
}

type application struct {
	dependenciesManager *DependenciesManager
}

func NewApplication(useDB bool) (Application, error) {
	return &application{dependenciesManager: NewDependenciesManager(useDB)}, nil
}

func (a *application) Run() error {
	router := gin.Default()
	a.initRouter(router)

	err := router.Run(":" + os.Getenv("SERVER_PORT"))

	return err
}

func (a *application) initRouter(router *gin.Engine) {
	hendler := a.dependenciesManager.linkPairHandler
	router.GET("/:shortLink", hendler.Get())
	router.POST("/", hendler.Post())
}
