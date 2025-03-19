package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	level := query.Get("level")

	switch level {
	case "info":
		infoLogger.Println("This is an info message")
	case "warn":
		warnLogger.Println("This is a warning message")
	case "debug":
		debugLogger.Println("This is a debug message")
	case "error":
		errorLogger.Println("This is an error message")
	default:
		infoLogger.Println("Default info log")
	}

	fmt.Fprintf(w, "Log generated with level: %s", level)
}

func main() {
	http.HandleFunc("/log", logHandler)

	infoLogger.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		errorLogger.Fatalf("Error starting server: %v\n", err)
	}
}
