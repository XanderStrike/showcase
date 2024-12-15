package main

import (
	"crypto/subtle"
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
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

func basicAuth(next http.HandlerFunc, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Skip auth if no username/password configured
		if username == "" || password == "" {
			next.ServeHTTP(w, r)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
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
	var config Conf
	confFile, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Printf("confFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(confFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	port := "8080"
	http.HandleFunc("/", basicAuth(HomeHandler, config.Username, config.Password))
	log.Printf("Server is ready to handle requests at port %s\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
