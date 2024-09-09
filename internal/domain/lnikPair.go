package domain

import "github.com/gin-gonic/gin"

type LinkPair struct {
	ShortLink string
	LongLink  string
}

type LinkPairRepository interface {
	Get(shortLink string) (LinkPair, error)
	Put(shortLing string, longLink string) error
	CheckLongLink(longLink string) (bool, LinkPair, error)
	CheckShortLink(longLink string) (bool, error)
}

type LinkPairService interface {
	Get(shortLink string) (LinkPair, error)
	Post(longLink string) (string, error)
}

type LinkPairHandler interface {
	Get() gin.HandlerFunc
	Post() gin.HandlerFunc
}
