package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

func main() {
	var conf GlobalConfiguration
	if _, err := toml.DecodeFile("varna.toml", &conf); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Start varna")

	router := NewRouter(&conf)

	log.Fatal(http.ListenAndServe(":8080", router))
}
