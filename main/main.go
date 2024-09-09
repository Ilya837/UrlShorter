package main

import (
	"Task1/internal/app"
	"flag"
	"github.com/joho/godotenv"
)

var useDB bool

func init() {
	flag.BoolVar(&useDB, "d", false, "Use data base.")
	err := godotenv.Load("task1.env")
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {

	flag.Parse()

	var application app.Application
	var err error

	if useDB {
		application, err = app.NewApplication(true)
	} else {
		application, err = app.NewApplication(false)
	}

	if err != nil {
		println(err)
		return
	}

	if err = application.Run(); err != nil {
		println(err)
		return
	}

}
