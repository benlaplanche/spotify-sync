package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server starting")

	api := RouterHandler(Router())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Port = ", port)

	http.ListenAndServe(":"+port, api)
}

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", RootPathHandler).Methods("GET")

	return router
}

func RouterHandler(router *mux.Router) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		router.ServeHTTP(res, req)
	}

}

func RootPathHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "hello world")
}
