package server

import (
	"log"
	handlers "myproject/internal/handlers"
	"net/http"
	"time"
)

/*
Поля настройки по умолчанию:

	Addr — используйте порт 8080.
	Handler — передайте ваш http-роутер.
	ErrorLog — передайте ваш логгер.
	ReadTimeout — таймаут для чтения. 5 секунд.
	WriteTimeout — таймаут для записи. 10 секунд.
	IdleTimeout — таймаут ожидания следующего запроса. 15 секунд.
*/

// 1.Создаем структуру сервера с полями для логгера (log.Logger) и http-сервера (http.Server).
type ServStruct struct {
	logger     *log.Logger
	HttpServer *http.Server
}

// --
// 2.Создаем функцию, в которой нужно создать http-роутер.
func Router(logger *log.Logger) *ServStruct {
	// Создаем роутер
	router := http.NewServeMux()

	// 3. Регистрируем хендлеры
	router.HandleFunc("/", handlers.HtmlHandler) // Для корневого пути
	router.HandleFunc("/upload", handlers.HandleUpload)

	// 4.Определяем сервер и передаем ему роутер
	srv := &http.Server{Addr: ":8080", Handler: router, ErrorLog: logger,
		ReadTimeout: 5 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 15 * time.Second}

	// 5.Возвращаем &ссылку на сервер
	return &ServStruct{logger, srv}
}
