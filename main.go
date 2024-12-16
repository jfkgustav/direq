package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	// "github.com/jfkgustav/direq/view/"
	"github.com/jfkgustav/direq/view/audience"
	"github.com/jfkgustav/direq/handler"
)

func main() {
	songs := handler.ReadRepertoireCSV()
	component := audience.Index(songs)
	
	http.Handle("/", templ.Handler(component))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}

