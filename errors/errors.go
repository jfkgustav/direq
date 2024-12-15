package errors

import "fmt"

var (
	MissingArtistErr = fmt.Errorf("Artist saknas")
	MissingSongErr   = fmt.Errorf("Låttitel saknas")
	MissingYearErr   = fmt.Errorf("År saknas")
)

// ["artist saknas", "år saknas"]
