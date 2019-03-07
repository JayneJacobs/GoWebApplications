package viewmodel

// MemberLocator type for page data
type MemberLocator struct {
	Title  string
	Active string
}

//NewMemberLocator adds the struct values
func NewMemberLocator() MemberLocator {
	result := MemberLocator{
		Active: "memberlocator",
		Title:  "Locate-Asylum Wind Band Members",
	}
	return result
}

// MemberCoordinate data items related to MemberLocation
type MemberCoordinate struct {
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}
