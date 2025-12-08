package main

import (
	"fmt"
	"net/http"
)

// logError() - универсальный вспомогательный инструмент для записи сообщений об ошибках.
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// ErrorResponse() является универсальным вспомогательным средством для отправки сообщений
// об ошибке в формате JSON с заданным кодом состояния. Обратите внимание, что мы используем интерфейс{}
// для параметра message, а не просто строковый тип. Это дает нам
// больше гибкости в выборе значений, которые мы можем включить в ответ.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// serverErrorResponse() будет использоваться, когда наше приложение столкнётся с
// непредвиденной проблемой во время выполнения. Он выводит подробное сообщение
// об ошибке, а затем с помощью вспомогательного метода errorResponse() отправляет
// клиенту код состояния 500 «Internal Server Error» и JSON-ответ (содержащий общее
// сообщение об ошибке).
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// notFoundResponse() используется для отправки клиенту кода состояния 404 "Not Found" и
// ответа в формате JSON.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// methodNotAllowedResponse() используется для отправки клиенту кода состояния 405 «Method Not Allowed»
// и ответа в формате JSON.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
