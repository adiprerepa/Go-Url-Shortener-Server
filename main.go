package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// populate with some urls
// we could use a persistent datastore in the future
var urlMap = map[string]string{
	"go": "https://golang.org",
	"fb" : "https://facebook.com",
	"cg" : "https://sfbay.craigslist.org",
	"nyt" : "https://www.nytimes.com",
	"ig" : "https://www.instagram.com",
}

// redirect to given url from {urlKey}
func redirectToUrl(w http.ResponseWriter, r *http.Request) {
	urlKey := mux.Vars(r)["urlKey"]
	fmt.Printf("UrlKey: %s\n", urlKey)
	if urlValue, ok := urlMap[urlKey]; ok {
		// getting value from map succeeded
		fmt.Printf("Redirecting to url %s...\n", urlValue)
		// redirect with code 303
		http.Redirect(w, r, urlValue, http.StatusSeeOther)
	} else {
		// 404 - no url in map
		_, _ = fmt.Fprintf(w, "404 url does not exist in map")
	}

}

// show all routes on home screen
func homeScreen(w http.ResponseWriter, r *http.Request) {
	jsonStr, _ := json.MarshalIndent(urlMap, "", "\t")
	initialUrls := "Welcome to the GCI Golang URL Shortener! Navigate to : GET /urls in order to access the urls. Available urls:\n"
	initialUrls += string(jsonStr)
	initialUrls += "\nTo get a URL, execute : GET /urls/{url_shortened}. For example, localhost:8080/urls/go would redirect to golang.org."
	_, _ = fmt.Fprintf(w, initialUrls)
}

func main() {
	fmt.Println("Starting Server...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeScreen).Methods("GET")
	router.HandleFunc("/urls/{urlKey}", redirectToUrl).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}