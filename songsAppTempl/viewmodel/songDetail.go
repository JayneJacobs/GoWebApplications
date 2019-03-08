package viewmodel

import "github.com/JayneJacobs/songsWebAppwtemplwDB/songsAppTempl/model"

//SongDetail Describes the Song page items
type SongDetail struct {
	Title  string
	Active string
	Music  []Music
}

// NewSongDetail provides the values for each song
func NewSongDetail(music []model.Music) SongDetail {
	result := SongDetail{
		Title:  "Asylum Wind Songs",
		Active: "song",
		Music:  []Music{},
	}
	for _, m := range music {
		result.Music = append(result.Music, musicToVM(&m))
	}
	return result
}
