package api

import (
	"log"
	"net/http"
)

func RunWebServer(port string) {
	log.Println("Start Http server at port " + port)
	router := NewRouter()
	http.Handle("/", router)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occurred starting HTTP listener at port" + port)
		log.Println("Error:", err.Error())
	}
}
