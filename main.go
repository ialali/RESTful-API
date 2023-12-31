package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"time"
)

var artists []Artist
var locations Locations
var dates Dates
var relations Relations
var client *http.Client

type Artist struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    Locations
	Dates        Dates
	Relations    Relations
}

type Locations struct {
	Index []Location `json:"index"`
}

type Location struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Index []Date `json:"index"`
}

type Date struct {
	ID    int64    `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID             int64                  `json:"id"`
	DatesLocations map[string]interface{} `json:"datesLocations"`
}

func Getjson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
func HandleHomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	data := map[string]interface{}{
		"Artists": artists,
	}
	// Serve the HTML page with the filtered artists
	tmpl := template.Must(template.ParseFiles("templates/template.html"))
	err := tmpl.Execute(w, data)
	if err != nil {
		return
	}
}
func AppendToStruct() {
	for index := range locations.Index {
		artists[index].Locations.Index = append(artists[index].Locations.Index, locations.Index[index])
	}

	for index := range dates.Index {
		artists[index].Dates.Index = append(artists[index].Dates.Index, dates.Index[index])
	}

	for index := range relations.Index {
		artists[index].Relations.Index = append(artists[index].Relations.Index, relations.Index[index])

	}
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	fs := http.FileServer(http.Dir("static"))

	Getjson("https://groupietrackers.herokuapp.com/api/artists", &artists)
	Getjson("https://groupietrackers.herokuapp.com/api/locations", &locations)
	Getjson("https://groupietrackers.herokuapp.com/api/dates", &dates)
	Getjson("https://groupietrackers.herokuapp.com/api/relations", &relations)
	AppendToStruct() // Associate locations, dates, and relations with each artist
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HandleHomePage)
	log.Println("Server listening http://localhost:8027")
	log.Fatal(http.ListenAndServe(":8027", nil))
}
