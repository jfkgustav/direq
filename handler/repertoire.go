package handler

import (
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"github.com/jfkgustav/direq/model"
	"github.com/jfkgustav/direq/errors"
)

func ReadRepertoireCSV() []*model.Song {
	file, err := os.Open("repertoire.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	var songs []*model.Song
	if unmarshalError := gocsv.UnmarshalFile(file, &songs); unmarshalError != nil {
		panic(unmarshalError)
	}
	for _, song := range songs {
		if song.Artist == "" {
			log.Println(errors.MissingArtistErr)
		}
	}
	return songs
}
