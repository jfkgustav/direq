package main

import (
	"github.com/gocarina/gocsv"
	"log"
	"os"
)

func ReadRepertoireCSV() []*Song {
	file, err := os.Open("repertoire.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	var songs []*Song
	if unmarshalError := gocsv.UnmarshalFile(file, &songs); unmarshalError != nil {
		panic(unmarshalError)
	}
	return songs
}
