package handlers

import (
	savedwords "langLearning/views/savedWords"
	"net/http"
)

func HandleServeSavedWordsPage(w http.ResponseWriter, r * http.Request){
	savedwords.SavedWordsPage().Render(r.Context(), w)
}
