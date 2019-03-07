package viewmodel

// Home Defines home page Title
type Home struct {
	Title  string
	Active string
}

// NewHome Defines home page Title
func NewHome() Home {
	result := Home{
		Active: "home",
		Title:  "AsylumWind Band",
	}
	return result
}
