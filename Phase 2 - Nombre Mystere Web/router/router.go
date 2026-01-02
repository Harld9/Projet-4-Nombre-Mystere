package router

import (
	"mystere/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/game", controller.Game)
	mux.HandleFunc("/end", controller.End)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mux
}
