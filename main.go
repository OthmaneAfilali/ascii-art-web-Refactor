package main

import (
	"ascii-art-web/pkg/generator"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type AsciiPage struct {
	Text   string
	Banner string
	Result string
	Error  string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", homeHandler)

	log.Println("Listening and serving on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homeHandler() handlers GET and POST request to "/" and "/ascii-art"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n\n=== Basic Request Info ===\n")
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL)
	fmt.Printf("Protocol: %s\n", r.Proto)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("Remote Address: %s\n", r.RemoteAddr)
	fmt.Printf("\n---------------------------")

	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		http.NotFound(w, r) // 404
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
		return
	}

	p := &AsciiPage{}
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK) // 200
		tmpl.Execute(w, p)

	case "POST":
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest) // 400
			return
		}
		text := r.Form.Get("text")
		banner := r.Form.Get("banner")
		if text == "" {
			p.Error = "Text input cannot be empty"
			tmpl.Execute(w, p)
			return
		}

		art, err := generator.GenArt(text, banner)
		if err != nil {
			p.Error = "Failed to generate ASCII art: " + err.Error()
			tmpl.Execute(w, p)
			return
		}

		// Populate the Page struct with form data and generated art
		p.Text = text
		p.Banner = banner
		p.Result = art

		w.WriteHeader(http.StatusOK) // 200
		tmpl.Execute(w, p)
	default:
		// Handle other methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // 405
	}

}
