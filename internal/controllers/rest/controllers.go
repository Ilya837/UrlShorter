package rest

import (
	"Task1/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostLinkPairRequest struct {
	LongLink string `json:"long_link"`
}

type GetLinkPairResponse struct {
	LongLink string `json:"long_link"`
}

type linkPairHandler struct {
	linkPairService domain.LinkPairService
}

func NewLinkPairHandler(linkPairService domain.LinkPairService) domain.LinkPairHandler {
	return &linkPairHandler{
		linkPairService: linkPairService,
	}
}
func (handler linkPairHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		shortLink := c.Params.ByName("shortLink")

		linkPair, err := handler.linkPairService.Get(shortLink)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, GetLinkPairResponse{linkPair.LongLink})

	}
}

func (handler linkPairHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postRequest PostLinkPairRequest

		err := c.BindJSON(&postRequest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res, err := handler.linkPairService.Post(postRequest.LongLink)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, res)

	}
}
