package viewmodel

import (
	"fmt"

	"github.com/JayneJacobs/songsWebAppwtemplAPI/songsAppTempl/model"
)

//Song definition Title Active and CAtegory Map
type Song struct {
	Title      string
	Active     string
	Categories []Category
}

// Category defines the List item
type Category struct {
	URL           string
	ImageURL      string
	Title         string
	Words         string
	IsOrientRight bool
	ChordsURL     string
}

// NewSongs is here
func NewSongs(categories []model.Category) Song {
	result := Song{
		Active: "songs",
		Title:  "Asylum Wind - Songs ",
	}
	result.Categories = make([]Category, len(categories))
	for i := 0; i < len(categories); i++ {
		vm := categorytoVM(categories[i])
		vm.IsOrientRight = i%2 == 1
		result.Categories[i] = vm
	}
	return result
}

func categorytoVM(c model.Category) Category {
	return Category{
		URL:       fmt.Sprintf("/songs/%v", c.ID),
		ImageURL:  c.ImageURL,
		Title:     c.Title,
		Words:     c.Words,
		ChordsURL: c.ChordsURL,
	}
}
