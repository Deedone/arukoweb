package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to ARUKO ORG</h1><p>Enjoy your stay</p>")
	fmt.Println("mainpage")
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.Execute(w, nil)
}

func submitRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	fmt.Fprintf(w, "nice")
}

func setupHTTP() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/submitRegister", submitRegister)
	http.HandleFunc("/register", registerPage)

	log.Fatal(http.ListenAndServe(":3030", nil))
}

func main() {
	setupHTTP()
}
