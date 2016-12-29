package main

import (
	"flag"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var GlobalConfig *Config

func main() {
	configFile := flag.String("config", "", "a configuration file to load")
	flag.Parse()
	loadConfiguration(*configFile)

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/results/:top", GetTopResults),
		rest.Post("/results", PostResult),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func GetTopResults(w rest.ResponseWriter, r *rest.Request) {
	tops := r.PathParam("top")
	if(tops == "") {
		// todo return bad request
	}

	top, err := strconv.Atoi(tops)
	if(err != nil) {
		// too return bad request
	}

	manager := DbManager{}
	scores := manager.Receive(top);
	w.WriteJson(&scores)
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

	manager := DbManager{}
	manager.Insert(&newResult)
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
