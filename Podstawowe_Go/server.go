package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Pobieranie portu z zmiennej środowiskowej lub ustawianie domyślnego portu na 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Pobranie strefy czasowej dla Polski
	serverLocation, err := time.LoadLocation("Local")
	if err != nil {
		log.Fatal("Error loading server location:", err)
	}

	// Logowanie informacji o dacie uruchomienia, autorze i porcie
	log.Printf("Server started at %s\n", time.Now().In(serverLocation).Format("2006-01-02 15:04:05"))
	log.Println("Author: Mikolaj Stasz")
	log.Println("Port:", port)

	// Obsługa żądania HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Pobranie adresu IP klienta
		ip := strings.Split(r.RemoteAddr, ":")[0]

		// Wyświetlenie danych na stronie
		fmt.Fprintf(w, "Client IP: %s\n", ip)
		fmt.Fprintf(w, "Server Time in Poland: %s\n", time.Now().In(serverLocation).Format("2006-01-02 15:04:05"))
	})

	// Uruchomienie serwera HTTP na określonym porcie
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
