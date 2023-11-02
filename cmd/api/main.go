package main

import (
	"fmt"
	"net/http"
	//"encoding/json"
	"time"
	"log"
	"os"
)

type CrawlerApp struct{
	log *log.Logger
	config AppConfig
}


func main(){
	fmt.Println("hello world")

	config := PopulateConfig()

	logger := log.New(os.Stdout,"INFO:",log.Llongfile|log.Ldate|log.Ltime)

	capp := &CrawlerApp{
	config : *config,
	log : logger,
	}

	log.Println("config loaded for address",capp.config.Address,"port",capp.config.Port )
	address := fmt.Sprintf("%s:%s",capp.config.Address,capp.config.Port)
	log.Println("address", address)
	// customer server
	server := &http.Server{
		Addr: ":8080",
		Handler: capp.routes(),
		IdleTimeout: time.Minute,
	}
	log.Println("server config",server.Addr)
	log.Fatal(server.ListenAndServe())
}

