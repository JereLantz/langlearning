package main

import (
	"encoding/json"
	"langLearning/handlers"
	"langLearning/views/home"
	"log"
	"net/http"
	"os"
	"strconv"
)

type configs struct {
	Port int `json:"port"`
}

var appConfigs configs

func parseConfigFile(path string) error{
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&appConfigs)
	if err != nil {
		return err
	}

	return nil
}

func init(){
	err := parseConfigFile("./.config.json")
	if err != nil {
		log.Fatalf("failed to parse the config file %s\n", err)
	}
}

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
		Addr: ":" + strconv.Itoa(appConfigs.Port),
		Handler: handler,
	}

	//pages
	handler.HandleFunc("GET /", handleServeHomepage)
	handler.HandleFunc("GET /newlesson", handlers.HandleServeNewLessonPage)

	log.Printf("Server started on port %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
