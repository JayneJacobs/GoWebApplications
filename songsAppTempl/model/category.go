package model

import (
	"errors"
	"fmt"
)

// Category struct is to populate the songs list
type Category struct {
	ID        int
	ImageURL  string
	Title     string
	Words     string
	ChordsURL string
}

var categories = []Category{
	Category{
		ID:        1,
		ImageURL:  "Girls.png",
		Title:     "Girls Just Want to have fun",
		Words:     `I come home in the morning Light`,
		ChordsURL: "https://tabs.ultimate-guitar.com/tab/cyndi_lauper/girls_just_want_to_have_fun_chords_1130196",
	}, Category{
		ID:       2,
		ImageURL: "kiwi.png",
		Title:    "Cups, Straws, and Other Supplies",
		Words: `From paper cups to bio-degradable plastic to straws and
						napkins, LSS is your source for the sundries that keep your stand
						running smoothly.`,
		ChordsURL: "https://tabs.ultimate-guitar.com/tab/cyndi_lauper/girls_just_want_to_have_fun_chords_1130196",
	}, Category{
		ID:       3,
		ImageURL: "pineapple.png",
		Title:    "Signs and Advertising",
		Words: `Sure, you could just wait for people to find your stand
						along the side of the road, but if you want to take it to the next
						level, our premium line of advertising supplies.`,
		ChordsURL: "https://tabs.ultimate-guitar.com/tab/cyndi_lauper/girls_just_want_to_have_fun_chords_1130196",
	},
}

// GetCategories returns categories
func GetCategories() []Category {
	return categories
}

// GetCategory returns  the Category or an error based on the ID
func GetCategory(ID int) (*Category, error) {
	for _, c := range categories {
		if c.ID == ID {
			return &c, nil
		}

	}
	return nil, errors.New(fmt.Sprintf("Category error for %v", ID))
}
