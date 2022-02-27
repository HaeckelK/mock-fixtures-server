package main

func pathToFileRouter(path string) (filename string, pattern string) {
	// Switch statement is evaluated in order.
	switch true {
	// Exact matches
	case matchRoute(path, "^/$"):
		filename = "/demo-fixtures/index.html"
		pattern = "^/$"
	case matchRoute(path, "^/person$"):
		filename = "/demo-fixtures/person-all.json"
		pattern = "^/person$"
	case matchRoute(path, `^/person/\d+$`):
		filename = "/demo-fixtures/person-specific.json"
		pattern = `^/person/\d+$`
	// Most specific non exact matches
	case matchRoute(path, `^/person/detail`):
		filename = "/demo-fixtures/person-all-detail.json"
		pattern = `^/person/detail`
	// Least specific non exact matches
	case matchRoute(path, "^/person"):
		filename = "/demo-fixtures/person-other.json"
		pattern = "^/person"
	default:
		filename = "/demo-fixtures/index.html"
		pattern = "NO MATCH"
	}
	return
}
