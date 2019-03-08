package controller

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/JayneJacobs/songsWebAppwtemplAPI/songsAppTempl/model"
	"github.com/JayneJacobs/songsWebAppwtemplAPI/songsAppTempl/viewmodel"
)

type song struct {
	songTemplate     *template.Template
	categoryTemplate *template.Template
	musicTemplate    *template.Template
}

func (s song) registerRoutes() {
	http.HandleFunc("/songs", s.handleSongs)
	http.HandleFunc("/songs/", s.handleSongs)
	http.HandleFunc("/music/", s.handleMusic)
}

func (s song) handleSongs(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/songs/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		s.handleCategory(w, r, categoryID)
	} else {

		categories := model.GetCategories()
		vm := viewmodel.NewSongs(categories)
		w.Header().Add("Content-Type", "text/html")
		s.songTemplate.Execute(w, vm)
	}
}

func (s song) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	music := model.GetMusicForCategory(categoryID)
	vm := viewmodel.NewSongDetail(music)
	w.Header().Add("Content-Type", "text/html")
	s.categoryTemplate.Execute(w, vm)
}

func (s song) handleMusic(w http.ResponseWriter, r *http.Request) {
	musicPattern, _ := regexp.Compile(`/music/(\d+)`)
	matches := musicPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		musicID, _ := strconv.Atoi(matches[1])
		music, err := model.GetMusic(musicID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		vm := viewmodel.NewMusic(music)
		s.musicTemplate.Execute(w, vm)
	} else {
		http.NotFound(w, r)
	}
}
