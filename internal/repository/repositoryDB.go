package repository

import (
	"Task1/internal/domain"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type LinkPairRepositoryDB struct {
	data *sql.DB
}

func (repo *LinkPairRepositoryDB) Get(shortLink string) (domain.LinkPair, error) {

	err := repo.data.Ping()

	if err != nil {
		return domain.LinkPair{}, err
	}

	rows, err := repo.data.Query("select long_link from public.link_pairs where short_link=$1", shortLink)

	if err != nil {
		return domain.LinkPair{}, err
	}

	defer rows.Close()

	if rows.Next() {
		var link string
		err = rows.Scan(&link)
		if err != nil {
			return domain.LinkPair{}, err
		}

		return domain.LinkPair{LongLink: link, ShortLink: shortLink}, nil
	} else {
		return domain.LinkPair{}, errors.New("link not found")
	}
}

func (repo *LinkPairRepositoryDB) Put(shortLink string, longLink string) error {

	err := repo.data.Ping()
	if err != nil {
		return err
	}

	query := "insert into public.link_pairs(short_link,long_link) values($1,$2)"
	_, err = repo.data.Exec(query, shortLink, longLink)

	if err != nil {
		return err
	}

	return nil
}

func (repo *LinkPairRepositoryDB) CheckLongLink(longLink string) (bool, domain.LinkPair, error) {

	query := `select short_link from public.link_pairs where long_link = $1`
	rows, err := repo.data.Query(query, longLink)

	if err != nil {
		return false, domain.LinkPair{}, err
	}

	defer rows.Close()

	if !rows.Next() {
		return false, domain.LinkPair{}, nil
	} else {
		var link string
		err = rows.Scan(&link)
		if err != nil {
			return false, domain.LinkPair{}, err
		}

		return true, domain.LinkPair{link, longLink}, nil

	}
}

func (repo *LinkPairRepositoryDB) CheckShortLink(shortLink string) (bool, error) {
	rows, err := repo.data.Query("select long_link from public.link_pairs where short_link = $1", shortLink)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	if !rows.Next() {
		return false, nil
	} else {
		return true, nil
	}

}

func NewLinkPairRepositoryDB(dbHost, dbPort, dbUser, dbPass, dbName string, createQuary string) domain.LinkPairRepository {

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	_, err = db.Query(createQuary)

	if err != nil {
		panic(err)
	}

	return &LinkPairRepositoryDB{
		data: db,
	}
}
