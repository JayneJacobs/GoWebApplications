package viewmodel

import (
	"html/template"

	"github.com/JayneJacobs/songsWebAppwtemplwDB/songsAppTempl/model"
)

//MusicViewModel defines Music Title
type MusicViewModel struct {
	Title  string
	Active string
	Music  Music
}

//Music defines Music
type Music struct {
	Name             string
	DescriptionShort string
	DescriptionLong  template.HTML
	PricePerSong     float32
	PricePerAlbum    float32
	Author           string
	IsOriginal       bool
	ImageURL         string
	ID               int
}

func musicToVM(music *model.Music) Music {
	return Music{
		Name:             music.Name,
		DescriptionShort: music.DescriptionShort,
		DescriptionLong:  template.HTML(music.DescriptionLong),
		PricePerSong:     music.PricePerSong,
		PricePerAlbum:    music.PricePerAlbum,
		Author:           music.Author,
		IsOriginal:       music.IsOriginal,
		ImageURL:         music.ImageURL,
		ID:               music.ID,
	}
}

// NewMusic describes a Music item (song)
func NewMusic(music *model.Music) MusicViewModel {
	return MusicViewModel{
		Title:  "Music by Asylum Wind",
		Active: "song",
		Music:  musicToVM(music),
	}
}
