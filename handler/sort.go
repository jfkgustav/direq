package handler

import (
	"github.com/jfkgustav/direq/model"
	"sort"
)

/*
func tags(s1, s2 *model.Song) bool {
	return s1.Tags[0] < s2.Tags[0]
}
*/

func SortSongs(songs []model.Song, sort_by string) []model.Song {
	switch sort_by {
	case "YearAsc":
		sort.Slice(songs, func(i, j int) bool {
			return songs[i].Year < songs[j].Year
		})
		break
	case "YearDes":
		sort.Slice(songs, func(i, j int) bool {
			return songs[i].Year > songs[j].Year
		})
		break
	case "ArtistAZ":
		sort.Slice(songs, func(i, j int) bool {
			return songs[i].Artist < songs[j].Artist
		})
		break
	case "ArtistZA":
		sort.Slice(songs, func(i, j int) bool {
			return songs[i].Artist > songs[j].Artist
		})
		break

	case "TitleAZ":
		sort.Slice(songs, func(i, j int) bool {
			return songs[i].Title > songs[j].Title
		})
		break
	case "TitleZA":
		sort.Slice(songs, func(i, j int) bool {
			return songs[i].Title > songs[j].Title
		})
		break
		/*
			case "TagAZ":
				break
			case "TagZA":
				break
		*/

	}
	return songs

}
