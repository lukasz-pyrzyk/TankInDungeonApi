package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"strconv"
)

func GetScoreResults(w rest.ResponseWriter, r *rest.Request) {
	ns := r.PathParam("n")
	if ns == "" {
		rest.Error(w, "N parameter is required", http.StatusBadRequest)
		return
	}

	top, err := strconv.Atoi(ns)
	if err != nil {
		rest.Error(w, "N parameter is invalid", http.StatusBadRequest)
		return
	}

	manager := NewDbManager()
	scores := manager.Receive(top, "scoreResults", "-score", "time");
	w.WriteJson(&scores)
}

func PostScoreResult(w rest.ResponseWriter, r *rest.Request) {
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
	manager.Insert(&newResult, "scoreResults")
	w.WriteHeader(http.StatusCreated)
}

func GetTimeResults(w rest.ResponseWriter, r *rest.Request) {
	ns := r.PathParam("n")
	if ns == "" {
		rest.Error(w, "N parameter is required", http.StatusBadRequest)
		return
	}

	top, err := strconv.Atoi(ns)
	if err != nil {
		rest.Error(w, "N parameter is invalid", http.StatusBadRequest)
		return
	}

	manager := NewDbManager()
	scores := manager.Receive(top, "timeResults", "-time", "-score");
	w.WriteJson(&scores)
}

func PostTimeResult(w rest.ResponseWriter, r *rest.Request) {
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
	manager.Insert(&newResult, "timeResults")
	w.WriteHeader(http.StatusCreated)
}