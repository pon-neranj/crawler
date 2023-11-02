package main

import(
	"net/http"
)

func (app *CrawlerApp)routes()(*http.ServeMux){
	server := http.NewServeMux()
	app.log.Println("in routes")
	server.HandleFunc("/health",app.healthHandler)
	server.HandleFunc("/startjob",app.startroutine)
	server.HandleFunc("/book",app.booklistHandler)
	server.HandleFunc("/book/",app.booksHandler)
	return server
}

