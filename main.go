package main

import (
	"langLearning/views/home"
	"log"
	"net/http"
)

func handleServeHomepage(w http.ResponseWriter, r *http.Request){
	home.Home().Render(r.Context(), w)
}

func main(){
	handler := http.NewServeMux()
	server := http.Server{
		Addr: ":42069",
		Handler: handler,
	}

	handler.HandleFunc("GET /", handleServeHomepage)
	log.Printf("Server started on port %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
