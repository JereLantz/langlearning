package handlers

import (
	newlesson "langLearning/views/newLesson"
	"net/http"
)

func HandleServeNewLessonPage(w http.ResponseWriter, r *http.Request){
	newlesson.NewLessonPage().Render(r.Context(), w)
}
