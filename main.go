package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var client *http.Client

type Artist struct {
	ID           int           `json:"id"`
	Image        string        `json:"image"`
	Name         string        `json:"name"`
	Members      []string      `json:"members"`
	Year         int           `json:"creationDate"`
	FirstAlbum   string        `json:"firstAlbum"`
	Locations    []Location    `json:"-"`
	ConcertDates []ConcertDate `json:"-"`
	Relations    []Relation    `json:"-"`
}

type Location struct {
	ID        int         `json:"id"`
	Locations []string    `json:"locations"`
	Date      ConcertDate `json:"-"`
}

type ConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Relation struct {
}

func GetArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []Artist
	err := GetJson(url, &artists)
	if err != nil {
		return nil, fmt.Errorf("error getting artist data: %s", err)
	}
	return artists, nil
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

func GetLocations() ([]string, error) {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	var location Location
	err := GetJson(url, &location)
	if err != nil {
		return nil, fmt.Errorf("error getting locations: %s", err)
	}
	return location.Locations, nil
}

func GetConcertDates() ([]string, error) {
	url := "https://groupietrackers.herokuapp.com/api/dates"

	var concertDates ConcertDate
	err := GetJson(url, &concertDates)
	if err != nil {
		return nil, fmt.Errorf("error getting concert dates: %s", err)
	}
	return concertDates.Dates, nil
}

func handleArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := GetArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Retrieved Artists:")
	for _, artist := range artists {
		fmt.Println(artist.Name)
	}

	tmpl := template.Must(template.ParseFiles("template.html"))
	err = tmpl.Execute(w, artists)
	if err != nil {
		return
	}
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	http.HandleFunc("/", handleArtists)

	log.Println("Server listening on http://localhost:8028")
	err := http.ListenAndServe(":8028", nil)
	if err != nil {
		log.Fatal(err)
	}
}
