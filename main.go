package main

import (
	"fmt"
	"net/http"

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
			tags := r.URL.Query()["tags"]
			fmt.Println(tags)
			audience.Index(songs, tags).Render(r.Context(), w)
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
