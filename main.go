package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jfkgustav/direq/handler"
	"github.com/jfkgustav/direq/model"
	"github.com/jfkgustav/direq/view/audience"
	"github.com/jfkgustav/direq/view/musician"
)

func main() {
	songs := handler.ReadRepertoireCSV()
	handler.SongRequests = make(map[int]model.SongRequest)

	http.HandleFunc("/mv", func(w http.ResponseWriter, r *http.Request) {
		sort := r.URL.Query().Get("sort")
		requestedSongs := handler.ReadRequests()

		fmt.Println("Requests ", len(requestedSongs))
		fmt.Printf("%v\n", requestedSongs)
		musician.Index(requestedSongs, sort == "pop").Render(r.Context(), w)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		showFilter := false
		tags := r.URL.Query()["tags"]
		decadeQuery := r.URL.Query().Get("decade")
		clear := r.URL.Query().Get("clear")
		fmt.Println(clear)
		var decade int
		decade, err := strconv.Atoi(decadeQuery)
		if err != nil || decadeQuery == "decade" {
			decade = 0
		}
		if len(tags) != 0 || decade != 0 || clear == "clear" {
			showFilter = true
		}
		// låtar = filter_funktion(filter_struct)
		audience.Index(songs, tags, decade, showFilter).Render(r.Context(), w)
	})

	http.HandleFunc("/request-song", func(w http.ResponseWriter, r *http.Request) {
		song_id_s := r.URL.Query().Get("song_id")
		var song_id int
		song_id, err := strconv.Atoi(song_id_s)
		if err != nil {
			w.Write([]byte(`Incorrect request`))
			return
		}
		handler.AddRequest(song_id)
		http.Redirect(w, r, "/", 200)
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
