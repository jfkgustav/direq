package handler

import (
	"github.com/gocarina/gocsv"
	"github.com/jfkgustav/direq/errors"
	"github.com/jfkgustav/direq/model"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

var Songs []model.Song
var Tags []string

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
		song.Title = songPre.Song
		song.Artist = songPre.Artist
		song.Year = songPre.Year
		song.Tags = strings.Split(songPre.Tags, ", ")
		Songs = append(Songs, song)
	}

	return Songs
}

func CreateTags() {
	for _, song := range Songs {
		for _, tag := range song.Tags {
			if !slices.Contains(Tags, tag) {
				Tags = append(Tags, tag)
			}
		}
	}
}

var SongRequests map[int]model.SongRequest

func RemoveRequest(song_id int) {
	delete(SongRequests, song_id)
}

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

// Bara gör saker på autopilot nu, och det här kanske är helt onödigt...
func SetSongs(songs []model.Song) {
	Songs = songs
}

func FilterSongs(tags []string, decade int) []model.Song {
	var filtered_songs []model.Song
	copy(filtered_songs, Songs)
	for _, song := range Songs {
		for _, tag := range song.Tags {
			if len(tags) == 0 || slices.Contains(tags, tag) {
				if decade == 0 || (song.Year >= decade && song.Year < decade+10) {
					filtered_songs = append(filtered_songs, song)
					break
				}
			}
		}
	}
	log.Println("Filtered", len(filtered_songs), "from total", len(Songs), "songs")
	return filtered_songs
}

var SongTags []string
