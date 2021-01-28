package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/register", http.StatusSeeOther)
	return
	fmt.Fprintf(w, "<h1>Welcome to ARUKO ORG</h1><p>Enjoy your stay</p>")
	fmt.Println("mainpage")
}

func registerPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.Execute(w, nil)
}

type regRequest struct {
	Name     string `json:"name"`
	Password string `json:"pass"`
	Confirm  string `json:"pass2"`
}

func submitRegister(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req regRequest
	err := decoder.Decode(&req)

	req.Name = strings.ToLower(req.Name)

	for _, char := range req.Name {
		if char < 'a' || char > 'z' {
			http.Error(w, "Name must not contain cpaces or non-letter", http.StatusBadRequest)
			return

		}
	}

	if err != nil {
		fmt.Println("Error decoding json")
		return
	}

	if req.Confirm != "tunnelsnakesrule" {
		http.Error(w, "Ask admins for secret password", http.StatusBadRequest)
		return
	}

	content, err := ioutil.ReadFile("/etc/passwd")
	etcpasswd := string(content)

	if len(req.Name) == 0 || len(req.Password) == 0 {
		http.Error(w, "Name and pass can't be empty", http.StatusBadRequest)
		return
	}
	if strings.Contains(etcpasswd, req.Name) {
		http.Error(w, "That user already exists", http.StatusBadRequest)
		return
	}

	cmd, err := exec.Command("./createuser.sh", req.Name, req.Password).Output()
	output := string(cmd)

	if err != nil {
		http.Error(w, output, http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintf(w, "Created user")
	}

}

func setupHTTP() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/submitRegister", submitRegister)
	http.HandleFunc("/register", registerPage)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":3030", nil))
}

func main() {
	setupHTTP()
}
