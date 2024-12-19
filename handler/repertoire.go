package handler

import (
	"github.com/gocarina/gocsv"
	"github.com/jfkgustav/direq/errors"
	"github.com/jfkgustav/direq/model"
	"log"
	"os"
	"strings"
	"time"
)

var Songs []model.Song

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

	for id, songPre := range songsPre {
		if songPre.Artist == "" {
			log.Println(errors.MissingArtistErr)
		}
		// if other errors....

		//finally
		var song model.Song
		song.ID = id
		song.Song = songPre.Song
		song.Artist = songPre.Artist
		song.Year = songPre.Year
		song.Tags = strings.Split(songPre.Tags, ", ")
		Songs = append(Songs, song)
	}

	return Songs
}

// funderar på om man kan ha en map istället?
// en map är per definition inte ordered, men eftersom varje request har en "created" så går det sortera visning efter den
// alternativt efter antal requests
var SongRequests map[int]model.SongRequest

func AddRequest(song_id int) {
	var song model.Song
	for _, s := range Songs {
		if s.ID == song_id {
			song = s
		}
	}
	log.Println("Add request for song", song)
	request, found := SongRequests[song_id]
	if found {
		request.NumberOfVotes++
		SongRequests[song_id] = request
	} else {
		song_request := model.SongRequest{
			Song:          song,
			Created:       time.Now(),
			NumberOfVotes: 1,
		}
		SongRequests[song_id] = song_request
	}

}

func ReadRequests() []model.SongRequest {
	songs := []model.SongRequest{}
	for _, r := range SongRequests {
		songs = append(songs, r)
	}
	return songs
}

/*
func ReadRequests() []model.SongRequest {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	file, err := os.Open("repertoire.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	var songsPre []*model.SongPre
	if unmarshalError := gocsv.UnmarshalFile(file, &songsPre); unmarshalError != nil {
		panic(unmarshalError)
	}

	var songs []model.SongRequest
		for _, songPre := range songsPre {
			if songPre.Artist == "" {
				log.Println(errors.MissingArtistErr)
			}
			// if other errors....

			//finally
			var request model.SongRequest
			var song model.Song
			song.Song = songPre.Song
			song.Artist = songPre.Artist
			song.Year = songPre.Year
			song.Tags = strings.Split(songPre.Tags, ", ")

			request.Song = song
			request.TimeInQueue = 1 + r.Intn(90)
			request.NumberOfVotes = 1 + r.Intn(4)
			songs = append(songs, request)
		}

	return songs
}
*/
