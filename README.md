# Go-Url-Shortener-Server
This is a tool for redirecting shortened urls. You can include a shortcut in a url, and it will redirect based on a pre-determined map.

## Setup
Execute `go get -u github.com/gorilla/mux` to install gorilla mux.

## Endpoints
- `GET /` - home, shows available urls.
- `GET /urls/{url}` - redirects to another url based on the shortened one you put in {url}

## Room For Improvement
- Use a persistent datastore for urls

## Author
- Aditya Prerepa. GCI 2020.
