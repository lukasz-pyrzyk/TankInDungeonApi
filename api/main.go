package main

import (
	"flag"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"strconv"
)

var DbHost *string

func main() {
	DbHost = flag.String("mongohost", "", "a mongodb host")
	flag.Parse()

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
	if tops == "" {
		rest.Error(w, "Top number is required", http.StatusBadRequest)
		return
	}

	top, err := strconv.Atoi(tops)
	if err != nil {
		rest.Error(w, "Top number is invalid", http.StatusBadRequest)
		return
	}

	manager := NewDbManager()
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

	manager := NewDbManager()
	manager.Insert(&newResult)
	w.WriteHeader(http.StatusCreated)
}

func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
