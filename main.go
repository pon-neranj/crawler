package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)
func main(){
	fmt.Println("hello world")

	server := http.NewServeMux()
	server.HandleFunc("/health",healthHandler)
	http.ListenAndServe(":8080", server)
}

func healthHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
	hReponse := map[string]string{
		"version":"1.0.0",
	}
	b,err := json.Marshal(hReponse)
	if err != nil {
		fmt.Println("ERROR")
	}
	w.Write(b)
}
