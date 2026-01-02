package router

import (
	"net/http"
	"quiz/controller"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.Home)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	mux.HandleFunc("/quizfacile", controller.QuizFacile)
	mux.HandleFunc("/quizmoyen", controller.QuizMoyen)
	mux.HandleFunc("/quizdifficile", controller.QuizDifficile)
	mux.HandleFunc("/reponse", controller.Reponse)

	return mux
}
