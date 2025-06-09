package handlers

import (
	"langLearning/views/lessons"
	"net/http"
)

func HandleServeLessonsPage(w http.ResponseWriter, r *http.Request){
	lessons.LessonsPage().Render(r.Context(), w)
}
