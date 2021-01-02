// Using Dependency Inversion to help test in Go
// This example shows how to perform test on http handler which calls a downstream
// service. The downstream service is mocked by a high-order-function constructor
// https://itnext.io/using-dependency-inversion-in-go-31d8bf9b3760
// https://medium.com/hatchpad/mocking-techniques-for-go-f0a302457d30

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuangwanjing/test_in_golang/http_downstream_mock/handlers"
)

func main() {
	r := mux.NewRouter()
	handler := handlers.NewHandlers(handlers.WithCustomerClient(&handlers.MockClient{}))
	r.HandleFunc("/", handler.MockHandler)
	log.Println("server running")
	log.Fatal(http.ListenAndServe(":8080", r))
}
