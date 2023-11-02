package main
import (
	"net/http"
	"webcrawler/internals/data"
	"encoding/json"
	"fmt"
)

func (app *CrawlerApp)healthHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.AppName)
	fmt.Fprintf(w, "version: %s\n", "v1.0.0")
}

func (app *CrawlerApp)startroutine(w http.ResponseWriter, r *http.Request){

}

func (app *CrawlerApp)booklistHandler(w http.ResponseWriter, r *http.Request){
	

	app.log.Println("IN booklistHandler")

	var book data.Book 
	newbook := book.FetchALL()
	b,err := json.Marshal(newbook)
	if err != nil {
		app.log.Println("error in marshal",err)
	}
	w.Write(b)
}

func (app *CrawlerApp)booksHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		var book data.Book
		newbook := book.Insert()
		b,err := json.Marshal(newbook)
		if err != nil {
			app.log.Println("error in marshal",err)
		}
		w.Write(b)
	}
}



