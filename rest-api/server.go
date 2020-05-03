package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter();
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	log.Println("Server listening on port", os.Getenv("PORT"))
	log.Fatalln(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}
