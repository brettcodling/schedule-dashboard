package scheduledashboard

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/brettcodling/jsonlogger"
	"github.com/brettcodling/scheduledashboard/pkg/utils"
	"github.com/go-co-op/gocron"
)

var Scheduler *gocron.Scheduler

// Start will start a new http server on the given port and then start the blocking scheduler
func Start(port string) {
	if _, err := os.Stat("favicon.ico"); errors.Is(err, os.ErrNotExist) {
		utils.CreateFavicon()
	}

	go func() {
		jsonlogger.Info("Starting")
		mux := http.NewServeMux()
		mux.HandleFunc("/favicon.ico", iconHandler)
		mux.HandleFunc("/dashboard", dashboardHandler)
		mux.HandleFunc("/jobs", jobHandler)
		jsonlogger.Info("listening")
		log.Fatal(http.ListenAndServe(":"+port, mux))
	}()

	Scheduler.StartBlocking()
}

// iconHandler will return the favicon
func iconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

// jobHandler will return the html for the jobs
func jobHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")

	for _, job := range Scheduler.Jobs() {
		w.Write([]byte(utils.BuildJobOutput(job)))
	}
}

// dashboardHandler will return the html for the dashboard layout
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "text/html")
	w.Write([]byte(utils.Header + utils.BodyStart + utils.BodyEnd + utils.Footer))
}
