package main

import (
	"ascii-art-web/pkg/generator"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Msg   string
}
type AsciiPg struct {
	P      Page
	Text   string
	Banner string
}

func main() {
	reqFiles := []string{
		"./assets/shadow.txt",
		"./assets/standard.txt",
		"./assets/thinkertoy.txt",
		"./static/styles.css",
		"./templates/index.html",
		"./templates/error.html",
	}

	checkRequired(reqFiles)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/", homeHandler)

	log.Println("Listening and serving on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homeHandler() handlers GET and POST request to "/" and "/ascii-art"
func homeHandler(w http.ResponseWriter, req *http.Request) {
	printRequest(req)
	if req.URL.Path != "/" && req.URL.Path != "/ascii-art" {
		http.Redirect(w, req, "/error?code=404", http.StatusNotFound)
		return
	}
	page := &AsciiPg{Banner: "standard"}
	if !isFileThere("./assets/" + page.Banner + ".txt") {
		http.Redirect(w, req, "/error?code=500", http.StatusInternalServerError)
		return
	}
	page.P.Title, _ = generator.GenArt("ASCII-ART", "standard")
	page.P.Msg, _ = generator.GenArt("Enter text here", "standard")

	switch req.Method {
	case "GET":
		getTemplate(w, req, "index", page)
	case "POST":
		handlePost(w, req, page)
	default:
		http.Redirect(w, req, "/error?code=405", http.StatusMethodNotAllowed)
	}
}

// getTemplate() grabs the template and
// write the response using the page info given
func getTemplate(w http.ResponseWriter, req *http.Request, tmplNm string, page any) {
	tmpl, err := template.ParseFiles("templates/" + tmplNm + ".html")
	if err != nil {
		http.Redirect(w, req, "/error?code=404", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	tmpl.Execute(w, page)
}

// handlePost() handles
func handlePost(w http.ResponseWriter, req *http.Request, page *AsciiPg) {
	text, banner, formErr := getFormInputs(req)
	if formErr != "" {
		http.Redirect(w, req, "/error?code=400", http.StatusBadRequest)
		return
	}
	if !isFileThere("./assets/" + banner + ".txt") {
		http.Redirect(w, req, "/error?code=500", http.StatusInternalServerError)
		return
	}
	if art, err := generator.GenArt(text, banner); err != nil {
		page.P.Msg = "Failed to generate ASCII art: " + err.Error()
		goto getTemplateLn
	} else {
		page.P.Msg = art
	}
	page.Text = text
	page.Banner = banner
getTemplateLn:
	getTemplate(w, req, "index", page)
}

// getFormInputs() gets the text and banner input from the form in the POST request
func getFormInputs(req *http.Request) (string, string, string) {
	if err := req.ParseForm(); err != nil {
		return "", "", "400"
	}
	text := req.Form.Get("text")
	banner := req.Form.Get("banner")

	if text == "" || banner == "" {
		return "", "Please type in your text and select banner style.", ""
	}
	return text, banner, ""
}

// errorHandler() generate custom error page responses.
// If error.html can't be parse, default to simple error page
func errorHandler(w http.ResponseWriter, req *http.Request) {
	statusCode := req.URL.Query().Get("code")
	page := &Page{Title: statusCode}
	switch statusCode {
	case "400":
		page.Msg = "400 bad request"
	case "404":
		page.Msg = "404 page not found"
	case "405":
		page.Msg = "405 method not allowed"
	case "500":
		page.Msg = "500 internal server error"
	default:
		page.Msg = "an unexpected error occurred"
	}
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
		return
	}
	page.Title, _ = generator.GenArt(page.Title, "standard")
	tmpl.Execute(w, page)
}

// checkRequired() checks a list of files if they exist or not.
// log.Fatal() if a file is not found
func checkRequired(reqFiles []string) {
	for _, file := range reqFiles {
		if !isFileThere(file) {
			log.Fatalf("Required file missing: %s\n", file)
		}
	}
}

// isFileThere() simply checks if a file exists
func isFileThere(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// printRequest() prints the first two lines of a HTTP Request header
func printRequest(r *http.Request) {
	log.Printf("\n=== HTTP Request Info ===\n"+
		"%s %s %s\n"+
		"Host: %s\n"+
		"---------------------------\n\n",
		r.Method, r.URL, r.Proto, r.Host)
}
