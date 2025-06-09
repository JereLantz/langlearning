package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"langLearning/handlers"
	"langLearning/views/home"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
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

func connectDb() (*sql.DB, error){
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initWordsTable(db *sql.DB) error{
	initWordsTableQuery := `CREATE TABLE IF NOT EXISTS words(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
		word TEXT UNIQUE NOT NULL,
		translation TEXT,
		learning_status INTERGER NOT NULL,
		language INTEGER NOT NULL,
		FOREIGN KEY(language) REFERENCES languages(id)
	);`

	_, err := db.Exec(initWordsTableQuery)
	if err != nil {
		return err
	}
	return nil
}

func initLessonsTable(db *sql.DB) error{
	initQuery := ` CREATE TABLE IF NOT EXISTS lessons(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
		language INTEGER,
		text BLOB NOT NULL,
		FOREIGN KEY(language) REFERENCES languages(id)
	);`

	_, err := db.Exec(initQuery)
	if err != nil {
		return err
	}
	return nil
}

func initLanguagesTable(db *sql.DB) error{
	initQuery := ` CREATE TABLE IF NOT EXISTS languages(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
		language TEXT UNIQUE NOT NULL
	);`

	_, err := db.Exec(initQuery)
	if err != nil {
		return err
	}
	return nil
}

func initDBSchema(db *sql.DB) error{
	err := initWordsTable(db)
	if err != nil{
		return fmt.Errorf("Initializing words table failed. %v", err)
	}

	err = initLanguagesTable(db)
	if err != nil{
		return fmt.Errorf("Initializing languages table failed. %v", err)
	}

	err = initLessonsTable(db)
	if err != nil{
		return fmt.Errorf("Initializing lessons table failed. %v", err)
	}

	return nil
}

func main(){
	handler := http.NewServeMux()
	server := http.Server{
		Addr: ":" + strconv.Itoa(appConfigs.Port),
		Handler: handler,
	}

	db, err := connectDb()
	if err != nil {
		log.Fatalf("Database connection could not be established %s\n", err)
	}

	err = initDBSchema(db)
	if err != nil {
		log.Fatalln(err)
	}

	//pages
	handler.HandleFunc("GET /", handleServeHomepage)
	handler.HandleFunc("GET /newlesson", handlers.HandleServeNewLessonPage)

	log.Printf("Server started on port %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
