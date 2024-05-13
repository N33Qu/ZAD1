package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"os"
)

func main() {
	author := "Mikolaj Stasz"

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverLocation, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal("Error loading server location:", err)
	}

	log.Printf("Server started at %s\n", time.Now().In(serverLocation).Format("2006-01-02 15:04:05"))
	log.Printf("Author: %s\n", author)
	log.Println("Port:", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/index.html")
	})

	http.HandleFunc("/api/time", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"time": time.Now().Format("2006-01-02 15:04:05")})
	})

	http.HandleFunc("/api/ip", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"ip": ip})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
