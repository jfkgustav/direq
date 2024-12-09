package main

import (
	"fmt"
)

func main() {
	songs := ReadRepertoireCSV()
	for _, song := range songs {
		fmt.Printf("%v\n", song)
	}
}
