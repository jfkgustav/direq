package model

import (
	"time"
)

// Song model as it is in base CSV
type SongPre struct {
	Status string `csv:"status"`
	Artist string `csv:"artist"`
	Song   string `csv:"song"`
	Year   int    `csv:"year"`
	Tags   string `csv:"tags"`
}

type Song struct {
	ID     int
	Status string
	Artist string
	Title  string
	Year   int
	Tags   []string
	// totaltVotes int
}

/*
type Vote struct {
	datetime   time.Time
	session_id int
}
*/

type SongRequest struct {
	Song          Song // kan bytas ut mot ID
	Created       time.Time
	NumberOfVotes int
}

type Session struct {
	ID             int
	Name           string
	SongsIDs       []int
	Date           time.Time
	PlayedSongsIDs []int
}
