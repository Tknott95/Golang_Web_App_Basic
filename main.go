package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type motiveQuote struct {
	Quote  string
	Author string
}

func init() {
	tpl = template.Must(template.ParseGlob("views/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// title := "Golang Admin Board"

	quote_1 := motiveQuote{
		Quote:  "If you ain't first you're last!",
		Author: "Ricky Bobby",
	}

	quote_2 := motiveQuote{
		Quote:  "Reach for the stars!",
		Author: "Buzz Lightyear",
	}

	both := []motiveQuote{quote_1, quote_2}

	tpl.ExecuteTemplate(w, "index.gohtml", both)
}

func about(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "about.gohtml", nil)
}

func contact(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "contact.gohtml", nil)
}
