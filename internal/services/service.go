package services

import (
	"Task1/internal/domain"
	"errors"
	"math/rand"
	"net/url"
)

type LinkPairService struct {
	linkPairRepository domain.LinkPairRepository
	serverLink         string
}

func NewLinkPairService(linkPairRepository domain.LinkPairRepository, host string, port string) LinkPairService {

	return LinkPairService{linkPairRepository: linkPairRepository, serverLink: "http://" + host + ":" + port + "/"}
}

func (linkPairService LinkPairService) Get(shortLink string) (domain.LinkPair, error) {
	result, err := linkPairService.linkPairRepository.Get(shortLink)

	if err != nil {
		return domain.LinkPair{}, err
	}

	return result, nil
}

func (linkPairService LinkPairService) Post(longLink string) (string, error) {

	u, err := url.Parse(linkPairService.serverLink)

	if err != nil {
		return "", err
	}

	if u.Scheme == "" || u.Host == "" {
		return "", errors.New("invalid url")
	}

	check, linkPair, err := linkPairService.linkPairRepository.CheckLongLink(longLink)
	if err != nil {
		return "", err
	}

	if check {

		return linkPairService.serverLink + linkPair.ShortLink, nil
	}

	var shortLink string

	for i := 0; i < 10; i++ {
		shortLink = CreateShortLink(7)
		tmp, err := linkPairService.linkPairRepository.CheckShortLink(shortLink)
		if err != nil {
			shortLink = ""
			continue
		}

		if !tmp {
			break
		} else {
			shortLink = ""
		}
	}

	if shortLink == "" {
		return "", errors.New("link does not created")
	}

	err = linkPairService.linkPairRepository.Put(shortLink, longLink)
	if err != nil {
		return "", err
	}
	return linkPairService.serverLink + shortLink, nil
}

func CreateShortLink(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
