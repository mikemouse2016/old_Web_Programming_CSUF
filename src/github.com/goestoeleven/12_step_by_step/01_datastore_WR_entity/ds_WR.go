package main

import (
	"fmt"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

type Employee struct {
	Name     string
	Role     string
	HireDate time.Time
	email    string
}

func init() {
	http.HandleFunc("/", myHandler)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	e1 := Employee{
		Name:     "David Jones",
		Role:     "Manager",
		HireDate: time.Now(),
		email:    "david@starship.com",
	}

	userKey := datastore.NewKey(c, "employee", e1.email, 0, nil)
	key, err := datastore.Put(c, userKey, &e1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var e2 Employee
	if err = datastore.Get(c, key, &e2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Wrote (Put) then read (Get) %q", e2.Name)
}
