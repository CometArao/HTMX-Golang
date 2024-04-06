package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Character struct {
	Id         int
	Name       string
	Race       string
	Image      string
	isSelected bool
}

var (
	characters = map[string][]Character{
		"Characters": {
			{Id: 1, Name: "Sergio Araya", Race: "Gato", Image: "/images/karin.png", isSelected: false},
			{Id: 2, Name: "Goku", Race: "Saiyajin", Image: "images/goku.png", isSelected: false},
			{Id: 3, Name: "Camara del tiempo", Race: "Templo", Image: "images/camara.jpg", isSelected: false},
			{Id: 4, Name: "Esfera", Race: "Esfera", Image: "images/esfera.jpg", isSelected: false},
			{Id: 5, Name: "Upa", Race: "Humano", Image: "images/upa2.png", isSelected: false},
			{Id: 6, Name: "Goten", Race: "50% Humano, 50% Saiyajin", Image: "images/goten.webp", isSelected: false},
			{Id: 7, Name: "Trunks", Race: "50% Humano, 50% Saiyajin", Image: "images/trunks.webp", isSelected: false},
		},
	}
	lastID = 7
)

func main() {
	fmt.Println("hello toriyama")

	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html")) // Revisa los templates disponibles

		tmpl.Execute(w, characters)
	}

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		// time.Sleep(2 * time.Second)
		name := r.PostFormValue("name")
		race := r.PostFormValue("race")
		image := r.PostFormValue("image")

		lastID++
		newCharacter := Character{Name: name, Race: race, Image: image, Id: lastID, isSelected: false}

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "character", newCharacter)

	}

	http.HandleFunc("/", handler1)
	http.HandleFunc("/add", handler2)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
