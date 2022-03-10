package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.gohtml")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.gohtml")
}
