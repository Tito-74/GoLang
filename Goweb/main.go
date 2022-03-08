package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)
type newsAggPage struct {
	Title string
	news string
}

func newAggHandler(w http.ResponseWriter, r *http.Request) {
	p:= newsAggPage{Title: "Amazing News Aggreator", news:"Great News for Africa"}
	t,_ :=template.ParseFiles("news.html")

	t.Execute(w, p)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
		http.HandleFunc("/news/", newAggHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}