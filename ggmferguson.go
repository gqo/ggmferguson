package main

import (
	"html/template"
	"log"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/books", bookHandler)
	http.HandleFunc("/album", albumHandler)

	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	http.ServeFile(w, r, "./web/static/main.html")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/images/favicon.ico")
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/static/books.html")
}

func albumHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./web/template/album.html"))
	Data := getAlbum()
	log.Println(Data)
	t.Execute(w, Data)
}
