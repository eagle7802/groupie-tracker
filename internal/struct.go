package internal

type Artist struct {
	ID             int      `json:"id"`
	Image          string   `json:"image"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	Locations      string   `json:"-"`
	ConcertDates   string   `json:"-"`
	DatesLocations map[string][]string
}

type Relation struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
type Search struct {
	Artists  []Artist
	Datalist map[string]string
}
