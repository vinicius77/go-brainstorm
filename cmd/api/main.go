package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/vinicius77/go-brainstorm/internal/handlers"
)

func main() {
	log.SetReportCaller(true)

	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting API service ...")
	fmt.Println("Server listening at http://localhost:8000 ...")

	err := http.ListenAndServe("localhost:8000", router)

	if err != nil {
		log.Error(err)
	}

}
