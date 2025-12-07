package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lekan-pvp/itbookworm/internal/data"
)

// Добавляем обработчик createBookHandler для эндпоинта `POST /v1/books`.
// Пока мы посто возвращаем текст.
func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create a new book")
}

// Добавляем обработчик showBookHandler для конечной точки `GET /v1/books/:id`.
// На данный момент мы получаем параметр id из URL-адреса и влючаем его в ответ.
func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		// используем новый метод notFoundResponse()
		app.notFoundResponse(w, r)
		return
	}

	book := data.Book{
		ID:       id,
		CreateAt: time.Now(),
		Title:    "Effective concurency in Go",
		Pages:    532,
		Genres:   []string{"IT"},
		Edition:  3,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		// используем новый метод для обработки ошибок сервера serverErrorResponse()
		app.serverErrorResponse(w, r, err)
	}
}
