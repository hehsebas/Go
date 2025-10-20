package main

import (
	"log"
	"net/http"
	"webappgo/handlers"
)

func main() {
	router := http.NewServeMux()
	//Static files
	fs := http.FileServer(http.Dir("static"))
	//Route to access static files
	router.Handle("/static/", http.StripPrefix("/static", fs))
	//Routes to handle the requests
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new-game", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/about", handlers.About)
	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", router)
}
