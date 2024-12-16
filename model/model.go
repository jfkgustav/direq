package model

// Song model as it is in base CSV
type SongPre struct {
	Status string `csv:"status"`
	Artist string `csv:"artist"`
	Song   string `csv:"song"`
	Year   int    `csv:"year"`
	Tags   string `csv:"tags"`
}

type Song struct {
	Status string 
	Artist string
	Song   string
	Year   int  
	Tags   []string
}

type songRequest struct {
	song 					Song
	timeInQueue 	int
	numberOfVotes int
}
