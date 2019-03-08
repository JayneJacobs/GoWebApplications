package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/JayneJacobs/songsWebAppwtemplwDB/songsAppTempl/viewmodel"
)

type memberLocator struct {
	memberLocatorTemplate *template.Template
}

func (ml memberLocator) registerRoutes() {
	http.HandleFunc("/member-locator", ml.handleMemberLocator)
	http.HandleFunc("/api/members", ml.handleAPIMembers)
}

//NewMemberLocator is the addresses of the band members houses
func (ml memberLocator) handleMemberLocator(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	vm := viewmodel.NewMemberLocator()
	ml.memberLocatorTemplate.Execute(w, vm)
}

func (ml memberLocator) handleAPIMembers(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var loc struct {
		ZipCode string `json:"zipCode"`
	}
	err := dec.Decode(&loc)
	if err != nil {
		log.Println(fmt.Errorf("Error retrievinglocation=: %v", err))
		enc := json.NewEncoder(w)
		enc.Encode([]viewmodel.MemberCoordinate{})
		return
	}
	log.Println("location:", loc)
	vm := coords
	enc := json.NewEncoder(w)
	enc.Encode(vm)
}

var coords []viewmodel.MemberCoordinate = []viewmodel.MemberCoordinate{
	viewmodel.MemberCoordinate{
		Latitude:  37.409,
		Longitude: -122.06,
		Title:     "Bobby's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.4092,
		Longitude: -122.061,
		Title:     "Macy's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.4094,
		Longitude: -122.06,
		Title:     "Juan's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.41,
		Longitude: -122.065,
		Title:     "Allison's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.415,
		Longitude: -122.07,
		Title:     "Chen's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.4217,
		Longitude: -122.075,
		Title:     "Matthew's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.4206,
		Longitude: -122.08,
		Title:     "Alice's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.4205,
		Longitude: -122.083,
		Title:     "Kara's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.414,
		Longitude: -122.09,
		Title:     "Fred's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.412,
		Longitude: -122.09,
		Title:     "Jake's Eatery location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.41,
		Longitude: -122.093,
		Title:     "Yaght Club's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.416,
		Longitude: -122.095,
		Title:     "Guitar Center's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.41,
		Longitude: -122.1,
		Title:     "Elvis's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  37.41,
		Longitude: -122.102,
		Title:     "Rich's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  40.412,
		Longitude: -75.099,
		Title:     "Scott's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  40.407,
		Longitude: -75.1025,
		Title:     "Jayne's location",
	},
	viewmodel.MemberCoordinate{
		Latitude:  40.191,
		Longitude: -75.026,
		Title:     "Jim's location",
	},
}
