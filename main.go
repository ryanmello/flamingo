package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error occuured", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "%s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request){
		log.Println("Goodbye world")
	})

	http.ListenAndServe(":9090", nil)
}