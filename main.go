package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

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
	var config Conf
	confFile, err := os.ReadFile("config/config.yml")
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
	port := "8080"
	http.HandleFunc("/", HomeHandler)
	log.Printf("Server is ready to handle requests at port %s\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
