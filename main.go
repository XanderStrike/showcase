package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	yaml "gopkg.in/yaml.v2"
)

type Tab struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type Conf struct {
	Title        string `yaml:"title"`
	Announcement string `yaml:"announcement"`
	Tabs         []Tab  `yaml:"tabs"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request!")
	var config Conf
	confFile, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Printf("confFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, config)
}

func main() {
	log.Println("Starting!")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
