package routes

import (
	"fmt"
	"net/http"
)

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server running!")
}

func healthPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Health fine")
}

func SetupRoutes() {
	http.HandleFunc("/", rootPathHandler)

	http.HandleFunc("/health", healthPathHandler)
}
