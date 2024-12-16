package handler

import (
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"strings"
	"github.com/jfkgustav/direq/model"
	"github.com/jfkgustav/direq/errors"
)

func ReadRepertoireCSV() []model.Song {
	file, err := os.Open("repertoire.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	var songsPre []*model.SongPre
	if unmarshalError := gocsv.UnmarshalFile(file, &songsPre); unmarshalError != nil {
		panic(unmarshalError)
	}

	var songs []model.Song
	for _, songPre := range songsPre {
		if songPre.Artist == "" {
			log.Println(errors.MissingArtistErr)
		}
		// if other errors....

		//finally
		var song model.Song
		song.Song = songPre.Song
		song.Artist = songPre.Artist
		song.Year = songPre.Year
		song.Tags = strings.Split(songPre.Tags, ", ")	
		songs = append(songs, song)
	}




	return songs
}
