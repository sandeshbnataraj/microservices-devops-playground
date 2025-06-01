package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const port = "8080"

func main() {
	log := logrus.New()
	log.Level = logrus.DebugLevel

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "🎉 Frontend is up — dummy mode")
	})

	addr := ":" + port
	log.Infof("starting dummy server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
