package models

type Artist struct{
	ID int `json:"id"`
	IMAGE string `json:"image"`
	NAME string `json:"name"`
	MEMBERS []string `json:"members"`
	CREATIONDATE int `json:"creationDate"`
	FIRSTALBUM string `json:"firstAlbum"`
	// "locations": "https://groupietrackers.herokuapp.com/api/locations/1",
    //  "concertDates": "https://groupietrackers.herokuapp.com/api/dates/1",
    //  "relations": "https://groupietrackers.herokuapp.com/api/relation/1"
}
type Locations struct{
	ID int `json:"id"`
	LOCATIONS []string `json:"locations"`
}
type ConcertDates struct{
	ID int `json:"id"`
	DATES []string `json:"dates"`
}
type Relations struct{
	ID int `json:"id"`
	IMAGE string `json:"image"`
	NAME string `json:"name"`
	MEMBERS []string `json:"members"`
	CREATIONDATE int `json:"creationDate"`
	FIRSTALBUM string `json:"firstAlbum"`

}
type LocationsWrapper struct {
    Index []Locations `json:"index"`
}

type DatesWrapper struct {
    Index []ConcertDates `json:"index"`
}

type RelationsWrapper struct {
    Index []Relations `json:"index"`
}

type PageData struct{
	Artist Artist
	Locations Locations
	ConcertDates ConcertDates
	Relations Relations
}