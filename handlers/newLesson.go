package handlers

import (
	"langLearning/views/components"
	"langLearning/views/newLesson"
	"net/http"
)

func HandleServeNewLessonPage(w http.ResponseWriter, r *http.Request){
	newlesson.NewLessonPage([]string{"sd", "sdfk"}).Render(r.Context(), w)
}

func HandleNewLangForm(w http.ResponseWriter, r *http.Request){
	components.NewLangForm().Render(r.Context(), w)
}
