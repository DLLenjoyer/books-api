package handlers

import (
	"encoding/json"
    "net/http"

    "github.com/DLLenjoyer/books-api/models"
    "github.com/DLLenjoyer/books-api/internal/service"
    "github.com/google/uuid"
    "github.com/gorilla/mux"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/books", h.getAllBooks).Methods(http.MethodGet)
	r.HandleFunc("/books/{id}", h.getBookByID).Methods(http.MethodGet)
	r.HandleFunc("/books", h.createBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{id}", h.updateBook).Methods(http.MethodPut)
	r.HandleFunc("/books", h.deleteBook).Methods(http.MethodDelete)
}

func (h *BookHandler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "Не удалось получить книги", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Не удалось конвертировать данные в файл JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *BookHandler) getBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	book, err := h.service.GetByID(bookID)
	if err != nil {
		http.Error(w, "Не удалось получить книгу по айди", http.StatusInternalServerError)
		return
	}
	if book == nil {
		http.Error(w, "Книга не найдена", http.StatusNotFound)
		return
	}

	jsonBytes, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Не удалось конвертировать данные в JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *BookHandler) createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Неправильный формат JSON", http.StatusInternalServerError)
		return
	}
	
	book.ID = uuid.New().String()

	if err := h.service.Add(&book); err != nil {
		http.Error(w, "Не удалось добавить книгу", http.StatusInternalServerError)
		return 
	}

	jsonBytes, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Не удалось конвертировать данные в JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBytes)

}

func (h *BookHandler) updateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Неправильный формат JSON", http.StatusBadRequest)
		return
	}

	if book.ID != bookID {
		http.Error(w, "ID книги не существует", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(&book); err != nil {
		http.Error(w, "Не удалось обновить книгу", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Не удалось конвертировать данные в JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *BookHandler) deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	if err := h.service.Delete(bookID); err != nil {
		http.Error(w, "Не удалось удалить книжку", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
