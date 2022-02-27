package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func matchRoute(path string, pattern string) bool {
	regex := regexp.MustCompile(pattern)
	matches := regex.FindStringSubmatch(path)

	if len(matches) <= 0 {
		return false
	}
	return true
}

type App struct {
	port                           string
	addPatternToResponseHeader     bool
	addFilenameToResponseHeader    bool
	addOriginalURLToResponseHeader bool
}

func (a App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.ToLower(r.URL.Path)
	// Trim trailing slashes
	if path[len(path)-1:] == "/" && len(path) > 1 {
		path = path[:len(path)-1]
	}

	filename, matchedPattern := pathToFileRouter(path)

	if a.addPatternToResponseHeader {
		w.Header().Add("x-matched-pattern", matchedPattern)
	}
	if a.addOriginalURLToResponseHeader {
		w.Header().Add("x-original-URL", r.URL.Path)
	}
	if a.addFilenameToResponseHeader {
		w.Header().Add("x-response-file", filename)
	}

	// Actually write the response
	http.ServeFile(w, r, filename)

}

func main() {
	// Load Environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := App{
		port:                           port,
		addPatternToResponseHeader:     true,
		addFilenameToResponseHeader:    true,
		addOriginalURLToResponseHeader: true,
	}

	fmt.Println("Starting fixtures server on", app.port)

	http.ListenAndServe(":"+app.port, app)
}
