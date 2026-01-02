package main

import (
	"fmt"
	"mystere/router"
	"net/http"
)

func main() {
	mux := router.New()
	fmt.Println("Serveur démarré sur http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Erreur serveur :", err)
	}
}
