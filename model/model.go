package model

type Song struct {
	Status string `csv:"status"`
	Artist string `csv:"artist"`
	Song   string `csv:"song"`
	Year   int    `csv:"year"`
	Tags   string `csv:"tags"`
}
