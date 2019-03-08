package model

import "fmt"

//Music is for the song purchase
type Music struct {
	Name             string
	DescriptionShort string
	DescriptionLong  string
	PricePerSong     float32
	PricePerAlbum    float32
	Author           string
	IsOriginal       bool
	ImageURL         string
	ID               int
	CategoryID       int
}

// GetMusicForCategory sets the Music items for Songs.html
func GetMusicForCategory(categoryID int) []Music {
	result := []Music{}
	for _, m := range music {
		if m.CategoryID == categoryID {
			result = append(result, m)
		}
	}
	return result
}

//GetMusic is for the list of songs
func GetMusic(musicID int) (*Music, error) {
	for _, m := range music {
		if m.ID == musicID {
			return &m, nil
		}
	}
	return nil, fmt.Errorf("Music not found with ID %v", musicID)
}

// Music defines the Music
var music = []Music{Music{
	Name:             "Girls Just Want to have Fun",
	DescriptionShort: "Made from fresh, organic California lemons.",
	DescriptionLong: `Made from premium, organic Meyer lemons. These fruit are left on the tree until they reach the peak of ripeness and then juiced within 8 hours of being picked.
			<br/>
			Lemonade made from our premium juice is sure to make your stand the most popular in the neighborhood.`,
	PricePerSong:  1.09,
	PricePerAlbum: 1.04,
	Author:        "California",
	IsOriginal:    true,
	ImageURL:      "Girls.png",
	ID:            1,
	CategoryID:    1,
}, Music{
	Name:             "Apple Juice",
	DescriptionShort: "The perfect blend of Washington apples.",
	DescriptionLong:  "The perfect blend of Washington apples.",
	PricePerSong:     0.89,
	PricePerAlbum:    0.85,
	Author:           "Ohio",
	IsOriginal:       true,
	ImageURL:         "apple.png",
	ID:               2,
	CategoryID:       1,
}, Music{
	Name:             "Watermelon Juice",
	DescriptionShort: "From sun-drenched fields in Florida.",
	DescriptionLong:  "From sun-drenched fields in Florida.",
	PricePerSong:     3.99,
	PricePerAlbum:    3.84,
	Author:           "Florida",
	IsOriginal:       true,
	ImageURL:         "watermelon.png",
	ID:               3,
	CategoryID:       1,
}, Music{
	Name:             "Kiwi Juice",
	DescriptionShort: "California sunshine and rain distilled into sweet goodness",
	DescriptionLong:  "California sunshine and rain distilled into sweet goodness",
	PricePerSong:     5.99,
	PricePerAlbum:    5.54,
	Author:           "California",
	IsOriginal:       false,
	ImageURL:         "kiwi.png",
	ID:               4,
	CategoryID:       1,
}, Music{
	Name:             "Mangosteen Juice",
	DescriptionShort: "Our latest taste sensation, imported directly from Hawaii",
	DescriptionLong:  "Our latest taste sensation, imported directly from Hawaii",
	PricePerSong:     6.87,
	PricePerAlbum:    6.79,
	Author:           "Hawaii",
	IsOriginal:       false,
	ImageURL:         "mangosteen.png",
	ID:               5,
	CategoryID:       1,
}, Music{
	Name:             "Orange Juice",
	DescriptionShort: "Fresh squeezed from Florida's best oranges.",
	DescriptionLong:  "Fresh squeezed from Florida's best oranges.",
	PricePerSong:     1.67,
	PricePerAlbum:    1.63,
	Author:           "Florida",
	IsOriginal:       false,
	ImageURL:         "orange.png",
	ID:               6,
	CategoryID:       1,
}, Music{
	Name:             "Pineapple Juice",
	DescriptionShort: "An exotic and refreshing offering. Straight from Hawaii.",
	DescriptionLong:  "An exotic and refreshing offering. Straight from Hawaii.",
	PricePerSong:     2.48,
	PricePerAlbum:    2.27,
	Author:           "Hawaii",
	IsOriginal:       false,
	ImageURL:         "pineapple.png",
	ID:               7,
	CategoryID:       1,
}, Music{
	Name:             "Strawberry Juice",
	DescriptionShort: "MThe perfect balance of sweet and tart.",
	DescriptionLong:  "The perfect balance of sweet and tart.",
	PricePerSong:     4.36,
	PricePerAlbum:    4.27,
	Author:           "California",
	IsOriginal:       false,
	ImageURL:         "strawberry.png",
	ID:               8,
	CategoryID:       1,
},
}
