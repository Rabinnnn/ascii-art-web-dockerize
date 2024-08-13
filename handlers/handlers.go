package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"

	asciiart "ascii-art-web/ascii-art"
)

// Index handles requests to "/" and "/Home"
func Index(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/", "/Home", "/home", "/HOME":
		if r.URL.Path == "/" || r.URL.Path == "/Home" || r.URL.Path == "/HOME" {
			http.Redirect(w, r, "/home", http.StatusMovedPermanently)
		}
		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/index.html")
		}
	case "/aboutus", "/Aboutus", "/ABOUTUS", "/ABOUT US", "/about us":
		if r.URL.Path == "/ABOUTUS" || r.URL.Path == "/Aboutus" || r.URL.Path == "about us" || r.URL.Path == "/ABOUT US" {
			http.Redirect(w, r, "/aboutus", http.StatusMovedPermanently)
		}
		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/aboutus.html")
		}
	case "/aboutascii", "/ascii", "/ABOUTASCII", "/About ascii", "/about ascii", "/ABOUT ASCII":
		if r.URL.Path == "/ABOUTASCII" || r.URL.Path == "/About ascii" || r.URL.Path == "about ascii" || r.URL.Path == "/ABOUT ASCII" {
			http.Redirect(w, r, "/aboutascii", http.StatusMovedPermanently)
		}
		if r.Method == http.MethodGet {
			// http.NotFound(w, r)
			serveTemplate(w, "templates/aboutascii.html")
		}
	case "/404":

		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/404.html")
		}
	case "/400", "/400.html":

		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/400.html")
			return
		}
	case "/500", "500.html":

		if r.Method == http.MethodGet {
			serveTemplate(w, "templates/500.html")
			return
		}

	default:
		if r.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("templates/404.html"))
			tmpl.Execute(w, "404: Page not Found!")

		}
	}
}

// HandleASCIIArt handles requests to "/ascii-art"
func HandleASCIIArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	bannerFile := r.FormValue("banner")
	input := r.FormValue("input")

	if len(input) == 0 || len(bannerFile) == 0 {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if containsIllegalCharacters(input) {
		http.Error(w, "Input contains illegal characters: Status 400: Bad Request", http.StatusBadRequest)
		return
	}

	bannerFilePath := asciiart.FileName(bannerFile)
	v := asciiart.CheckStatus(bannerFilePath)
	if v != nil {
		http.NotFound(w, r)
		return
	}
	_, errr := os.Stat("templates/result.html")
	if os.IsNotExist(errr) {
		http.NotFound(w, r)
		return
	}

	tmpl, eE := template.ParseFiles("templates/result.html")
	if eE != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("500 Internal Server Error")
		return
	}
	ascii := asciiart.Art(input, asciiart.BannerArt(bannerFilePath))
	e := tmpl.Execute(w, ascii.String())
	if e != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("500 Internal Server Error")
		return
	}
}

// serveTemplate loads and executes a template file
func serveTemplate(w http.ResponseWriter, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		http.Error(w, "404 Page not Found", http.StatusNotFound)
		return
	}
	tmpl, ee := template.ParseFiles(filename)
	if ee != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("500 Internal Server Error")
		return
	}
	errr := tmpl.Execute(w, nil)
	if errr != nil {
		log.Println("500 Internal Server Error")
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Function to check for illegal characters in input
func containsIllegalCharacters(input string) bool {
	// Regular expression to match non-printable ASCII characters
	illegalCharRegex := regexp.MustCompile(`[^\x00-\x7F]`)
	return illegalCharRegex.MatchString(input)
}
