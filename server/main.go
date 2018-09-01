package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type configuration struct {
	Info string `json:"info"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}

type response struct {
	ClientInfo string    `json:"clientInfo"`
	Info       string    `json:"info"`
	Time       time.Time `json:"time"`
}

var config configuration

func greet(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request...")
	keys := r.URL.Query()

	data := response{
		ClientInfo: keys["info"][0],
		Info:       config.Info,
		Time:       time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
	log.Println("Response sent")
}

func main() {
	c := flag.String("c", "./config.json", "json config")
	flag.Parse()
	file, err := os.Open(*c)

	if err != nil {
		log.Fatal("cant open config file", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}

	log.Printf("Listening on %s", config.Port)
	http.HandleFunc("/", greet)
	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil)
}
