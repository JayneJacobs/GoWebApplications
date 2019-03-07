package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController          home
	songController          song
	memberLocatorController memberLocator
)

// Startup is the Controller entry point
func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	songController.songTemplate = templates["songs.html"]
	homeController.loginTemplate = templates["login.html"]
	songController.categoryTemplate = templates["song_details.html"]
	songController.musicTemplate = templates["song_detail.html"]
	memberLocatorController.memberLocatorTemplate = templates["member_locator.html"]
	homeController.registerRoutes()
	songController.registerRoutes()
	memberLocatorController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
