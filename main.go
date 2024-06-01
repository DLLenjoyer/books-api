package main

import (
	"log"
	"net/http"

	"github.com/DLLenjoyer/books-api/internal/handlers"
	"github.com/DLLenjoyer/books-api/internal/repository"
	"github.com/DLLenjoyer/books-api/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewInMemoryBook()
	bookService := service.NewBookService(repo)
	bookHandler := handlers.NewBookHandler(bookService)

	r := mux.NewRouter()
	bookHandler.SetupRoutes(r)

	log.Println("starting server on :6030")
	log.Fatal(http.ListenAndServe(":6030", r))
}
