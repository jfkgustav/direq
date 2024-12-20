package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jfkgustav/direq/handler"
	"github.com/jfkgustav/direq/model"
	"github.com/jfkgustav/direq/store"
	"github.com/jfkgustav/direq/view/audience"
	"github.com/jfkgustav/direq/view/backstage"
	"github.com/jfkgustav/direq/view/backstage/session"
	"github.com/jfkgustav/direq/view/musician"
)

func main() {
	db := store.NewBoltDB("./direq.db")
	defer db.Close()
	songs := handler.ReadRepertoireCSV()
	err := store.UpdateRepertoire(db, songs)

	if err != nil {
		log.Fatal(err)
	}

	// autopilothjärnan kickade in här!

	songs, err = store.GetSongs(db)
	if err != nil {
		log.Fatal(err)
	}
	handler.SetSongs(songs)

	handler.CreateTags()

	handler.SongRequests = make(map[int]model.SongRequest)

	http.HandleFunc("/mv", func(w http.ResponseWriter, r *http.Request) {
		sort := r.URL.Query().Get("sort")
		requestedSongs := handler.ReadRequests()

		fmt.Println("Requests ", len(requestedSongs))
		fmt.Printf("%v\n", requestedSongs)
		musician.Index(requestedSongs, sort == "pop").Render(r.Context(), w)
	})

	http.HandleFunc("/backstage", func(w http.ResponseWriter, r *http.Request) {
		backstage.Index().Render(r.Context(), w)
	})

	http.HandleFunc("/new-session", func(w http.ResponseWriter, r *http.Request) {
		session.NewSession().Render(r.Context(), w)
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
		sort := r.URL.Query().Get("sort")
		// låtar = filter_funktion(filter_struct)
		filtered_songs := handler.FilterSongs(tags, decade)

		sorted_songs := handler.SortSongs(filtered_songs, sort)
		audience.Index(sorted_songs, tags, decade, showFilter, sort).Render(r.Context(), w)
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

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
