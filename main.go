package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jfkgustav/direq/view"
	"github.com/jfkgustav/direq/handler"
)

func main() {
	component := view.Hello("John")
	
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
	songs := handler.ReadRepertoireCSV()
	for _, song := range songs {
		fmt.Printf("%v\n", song)
	}
}

