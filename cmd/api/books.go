package main

import (
	"fmt"
	"net/http"
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
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of book %d\n", id)
}
