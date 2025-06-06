package main

import (
	"langLearning/handlers"
	"langLearning/views/home"
	"log"
	"net/http"
)

func handleServeHomepage(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		w.WriteHeader(404)
		return
	}
	home.Home().Render(r.Context(), w)
}

func main(){
	handler := http.NewServeMux()
	server := http.Server{
		Addr: ":42069",
		Handler: handler,
	}

	//pages
	handler.HandleFunc("GET /", handleServeHomepage)
	handler.HandleFunc("GET /newlesson", handlers.HandleServeNewLessonPage)

	log.Printf("Server started on port %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
