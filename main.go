package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jfkgustav/direq/view/musician"
	"github.com/jfkgustav/direq/view/audience"
	"github.com/jfkgustav/direq/handler"
)

func main() {
	songs := handler.ReadRepertoireCSV()
	requestedSongs := handler.ReadRequests();


	http.HandleFunc("/mv", func(w http.ResponseWriter, r *http.Request) {
			sort := r.URL.Query().Get("sort")
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
			if err != nil || decadeQuery == "decade"{
				decade = 0
			}
			if len(tags) != 0 || decade != 0 || clear == "clear" {
				showFilter = true
			}

			audience.Index(songs, tags, decade, showFilter).Render(r.Context(), w)
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
