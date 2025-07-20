package main

import (
	"log"
	server "myproject/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(
		log.Writer(),
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	// Создаем сервер
	srv := server.Router(logger)
	err := srv.HttpServer.ListenAndServe()
	if err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}

}
