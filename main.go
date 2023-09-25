package main

import (
	"GoProjects/database"
	"GoProjects/routes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer db.Close()

	router := routes.SetupRouter()

	router.Use(loggingMiddleware)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a JSON log entry for the request
		logEntry := map[string]interface{}{
			"method": r.Method,
			"path":   r.URL.Path,
			"remote": r.RemoteAddr,
		}

		// Marshal the log entry to JSON
		logData, err := json.Marshal(logEntry)
		if err != nil {
			log.Printf("Error marshaling log entry: %v", err)
		} else {
			log.Println(string(logData))
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
