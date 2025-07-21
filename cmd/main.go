package main

import (
	"log"
	"os"

	server "myproject/internal/server"
)

func main() {
	// Создаем логгер
	//// Файл для логов
	logFile, err := os.OpenFile("logmorse.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()

	logger := log.New(
		logFile, //log.Writer(),
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	// Создаем сервер
	srv := server.Router(logger)
	err = srv.HttpServer.ListenAndServe()
	if err != nil {
		logger.Fatalf("Ошибка при запуске сервера: %v", err)
	}

}
