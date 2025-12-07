package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	// Инициализируем новый маршрутизатор chi
	router := chi.NewRouter()

	// Преобразуем вспомогательную функцию notFoundResponse() в http.Handler, а затем устанавливаем его
	// в качестве пользовательского обработчика ошибок для ответов 404 Not Found
	router.NotFound(app.notFoundResponse)

	// Аналогичным образом преобразуем вспомогательную функцию methodNotAllowedResponse() в http.Handler и установите
	// её в качестве пользовательского обработчика ошибок для ответов 405 Method Not Allowed
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	// Регистрируем шаблоны URL и методы обработчиков
	router.Get("/v1/healthcheck", app.healthcheckHandler)
	router.Post("/v1/books", app.createBookHandler)
	router.Get("/v1/books/{id}", app.showBookHandler)

	// Возвращаем экземпляр chi.Mux
	return router
}
