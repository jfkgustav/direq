package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jfkgustav/direq/handler"
	"github.com/jfkgustav/direq/model"
	"github.com/jfkgustav/direq/view/audience"
	"github.com/jfkgustav/direq/view/musician"
	"github.com/jfkgustav/direq/view/backstage"
	"github.com/a-h/templ"	
)

func main() {
	handler.ReadRepertoireCSV()
	handler.CreateTags()
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
		filtered_songs := handler.FilterSongs(tags, decade)
		audience.Index(filtered_songs, tags, decade, showFilter).Render(r.Context(), w)
	})

	http.HandleFunc("/request-song", func(w http.ResponseWriter, r *http.Request) {
		song_id_s := r.URL.Query().Get("song_id")
		var song_id int
		song_id, err := strconv.Atoi(song_id_s)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`Invalid request, should be number`))
			return
		}
		handler.AddRequest(song_id)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	})

	backstage := backstage.Index()
	http.Handle("/backstage", templ.Handler(backstage))


	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
