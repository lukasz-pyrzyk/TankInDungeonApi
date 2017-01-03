package main

import (
	"flag"
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

var DbHost *string

func main() {
	DbHost = flag.String("mongohost", "", "a mongodb host")
	flag.Parse()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/scoreresults/:n", GetScoreResults),
		rest.Post("/scoreresults", PostScoreResult),
		rest.Get("/timeresults/:n", GetTimeResults),
		rest.Post("/timeresults", PostTimeResult),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func failOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
