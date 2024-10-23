package funcs

import (
	"encoding/json"
	"net/http"
)


// This function fetches a list of artists from the API and stores them in the Artists slice.
func fetchArtists() {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		Artists = append(Artists, ArtistData{FEATCHINGerror: err})
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&Artists)
	if err != nil {
		Artists = append(Artists, ArtistData{FEATCHINGerror: err})
		return
	}
}

// A method on the ArtistData struct that fetches the relations of an artist from their Relations link.
func (a *ArtistData) fetchRelations() {
	resp, err := http.Get(a.Relations)
	if err != nil {
		a.FEATCHINGerror = err
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&a.RelationData)
	if err != nil {
		a.FEATCHINGerror = err
		return
	}
}

// A method that fetches the artist's locations data.
func (a *ArtistData) fetchLocation() {
	resp, err := http.Get(a.Locations)
	if err != nil {
		a.FEATCHINGerror = err
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&a.LocationsDATA)
	if err != nil {
		a.FEATCHINGerror = err
		return
	}
}

// A method that fetches the concert dates of the artist.
func (a *ArtistData) fetchDates() {
	resp, err := http.Get(a.ConcertDates)
	if err != nil {
		a.FEATCHINGerror = err
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&a.ConcertDatesDATA)
	if err != nil {
		a.FEATCHINGerror = err
		return
	}
}

// This is a wrapper method that calls the other fetch methods (fetchRelations, fetchDates, and fetchLocation) for an artist.
func (a *ArtistData) FeatchAll() {
	a.fetchRelations()
	a.fetchDates()
	a.fetchLocation()
}



