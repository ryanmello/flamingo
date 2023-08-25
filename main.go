package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request){
		log.Println("Goodbye world")
	})

	http.ListenAndServe(":9090", nil)
}