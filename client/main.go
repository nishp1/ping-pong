package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type configuration struct {
	Source source `json`
	Target target `json`
}

type source struct {
	Info string `json`
	IP   string `json`
}

type target struct {
	Info     string `json`
	Protocol string `json`
	IP       string `json`
	Port     string `json`
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }

func main() {
	c := flag.String("c", "./config.json", "json config")
	flag.Parse()
	file, err := os.Open(*c)

	if err != nil {
		log.Fatal("cant open config file", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("can't decode config JSON: ", err)
	}

	v := url.Values{"info": {config.Source.Info}}
	url := fmt.Sprintf("http://%s:%s?%s", config.Target.IP, config.Target.Port, v.Encode())

	response, err := http.Get(url)

	if err != nil {
		log.Fatalf("%s", err)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("%s\n", string(contents))

}
