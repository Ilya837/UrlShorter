package repository

import (
	"Task1/internal/domain"
	"errors"
)

type LinkPairRepository struct {
	data map[string]string
}

func (repo *LinkPairRepository) Get(shortLink string) (domain.LinkPair, error) {
	res := repo.data[shortLink]
	println(shortLink)
	if res == "" {
		err := errors.New("short link not found")
		return domain.LinkPair{}, err
	}

	return domain.LinkPair{ShortLink: shortLink, LongLink: res}, nil

}

func (repo *LinkPairRepository) Put(shortLink string, longLink string) error {

	repo.data[shortLink] = longLink
	return nil
}

func (repo *LinkPairRepository) CheckLongLink(longLink string) (bool, domain.LinkPair, error) {
	for k, v := range repo.data {
		if v == longLink {
			return true, domain.LinkPair{ShortLink: k, LongLink: v}, nil
		}
	}
	return false, domain.LinkPair{}, nil
}

func (repo *LinkPairRepository) CheckShortLink(shortLink string) (bool, error) {
	if repo.data[shortLink] == "" {
		return false, nil
	}
	return true, nil
}

func NewLinkPairRepository() domain.LinkPairRepository {
	return &LinkPairRepository{
		data: make(map[string]string, 1024),
	}
}
