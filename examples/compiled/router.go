// Write your own custom router to match URLs to filenames
package main

func pathToFileRouter(path string) (filename string, pattern string) {
	// Switch statement is evaluated in order.
	switch true {
	// Exact matches
	case matchRoute(path, "^/$"):
		filename = "/fixtures/index.html"
		pattern = "^/$"
	case matchRoute(path, "^/colour$"):
		filename = "/fixtures/colours-all.json"
		pattern = "^/colour$"
	default:
		filename = "/fixtures/index.html"
		pattern = "NO MATCH"
	}
	return
}
