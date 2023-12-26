package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jacobslunga/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
  log.SetReportCaller(true)
  var r *chi.Router = chi.NewRouter()
  handlers.Handler(r)

  fmt.Println("Starting our Go API")

  err := http.ListenAndServe("localhost:8000", r)

  if err != nil {
    log.Fatal(err)
  }
}
