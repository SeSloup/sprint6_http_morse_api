package handlers

import (
	"fmt"
	"io"
	service "myproject/internal/service"
	"net/http"
	"os"
)

// Хендлер для корневого эндпоинта
func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	// Читаем содержимое файла index.html
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Отправляем содержимое файла в ответ
	w.Write(data)
}

// Хендлер для эндпоинта /upload
func HandleUpload(w http.ResponseWriter, r *http.Request) {

	// 1.Парсим html-форму из файла index.html.
	err := r.ParseMultipartForm(10) // ограничение на 10 МБ
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusBadRequest)
		return
	}
	fmt.Println("file upload") //заменить на log

	// 2.Получаем файл из формы
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusBadRequest)
		return
	}
	defer file.Close() // закрываем файл по окончании работы функции

	// 3.Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	// 4. Передаем данные в функцию из пакета service для конвертации
	textForm := string(data)
	conv, isMorse := service.TestPrint(textForm)

	// 5-7. Создаем локальный файл и записываем в него результат строки. Возвращаем результат конвертации

	if isMorse {
		//fmt.Fprintf(w, "Это расшифрованный текст: %s", conv)
		fmt.Fprintf(w, "%s", conv)
		service.CreateFile("text.txt", conv, "./output")
	} else {
		//fmt.Fprintf(w, "Это зашифрованный текст: %s", conv)
		fmt.Fprintf(w, "%s", conv)
		service.CreateFile("morse.txt", conv, "./output")
	}

	// Отправляем содержимое файла в ответ
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)

	/*s := fmt.Sprintf("\nMethod: %s\nHost: %s\nPath: %s",
		r.Method, r.Host, r.URL.Path)
	w.Write([]byte(s))
	*/
}
