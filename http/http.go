package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Person person
type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// IndexHandler index handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world hoge")
}

var t = template.Must(template.ParseFiles("index.html"))

// PersonHandler person handler
func PersonHandler(w http.ResponseWriter,
	r *http.Request) {
	defer r.Body.Close()

	if r.Method == "POST" {
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		person := Person{
			ID:   id,
			Name: string(b),
		}

		t.Execute(w, person)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":3000", nil)
}
