package main

import (
	"flag"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

var GlobalConfig *Config
var store []Result

func main() {
	configFile := flag.String("config", "", "a configuration file to load")
	flag.Parse()
	loadConfiguration(*configFile)

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/results", GetTopResults),
		rest.Post("/results", PostResult),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))

	store = make([]Result, 10, 20)
}

func GetTopResults(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(&store)
}

func PostResult(w rest.ResponseWriter, r *rest.Request) {
	newResult := Result{}
	err := r.DecodeJsonPayload(&newResult)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = newResult.Validate()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store = append(store, newResult)
	w.WriteHeader(http.StatusCreated)
}

func loadConfiguration(cfgFile string) {
	data, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		e := err.Error()
		panic(fmt.Sprintf("%s: %s", e, err))
	}

	yaml.Unmarshal(data, &GlobalConfig)
}

func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
