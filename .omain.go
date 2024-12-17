package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jfkgustav/direq/view/musician"
	"github.com/jfkgustav/direq/view/audience"
	"github.com/jfkgustav/direq/handler"
)

func main() {
	songs := handler.ReadRepertoireCSV()
	index := audience.Index(songs)
	requestedSongs := handler.ReadRequests();
	musiciansView := musician.Index(requestedSongs);

	
	http.Handle("/", templ.Handler(index))
	http.Handle("/mv", templ.Handler(musiciansView));
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

