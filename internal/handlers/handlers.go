package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	service "myproject/internal/service"
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
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Ошибка записи содержимого файла в ответ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Хендлер для эндпоинта /upload
func HandleUpload(w http.ResponseWriter, r *http.Request, logger *log.Logger) {

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
	conv, err := service.TestPrint(textForm)
	if err != nil {
		http.Error(w, "Ошибка конвертации файла морзе. Передана пустая строка", http.StatusInternalServerError)
		return
	}

	// 5-7. Создаем локальный файл и записываем в него результат строки. Возвращаем результат конвертации
	// Отправляем содержимое файла в ответ
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fmt.Fprintf(w, "%s", conv) //Задача Вернуть результат конвертации строки. Куда? - конкретики нет. Поэтому возвращаем в ответ пользователю, в файл и в лог
	logger.Printf("Текст конвертации: %s", conv)

	name := filepath.Base("text.txt")
	ext := filepath.Ext("text.txt")
	time := time.Now().UTC().Format("20060102_150405")
	filename := fmt.Sprint(name[:len(name)-len(ext)], "_", time, ext)

	// формируем относительное имя нужного файла
	dir := "./output"
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	filename = filepath.Join(dir, filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// сохраняем идентификатор текущего вывода
	stdout := os.Stdout
	// присваиваем os.Stdout идентификатор открытого файла
	os.Stdout = f
	// строка должна записаться в файл
	fmt.Println(conv)

	// возвращаем обратно вывод в консоль
	os.Stdout = stdout
	// строка выведется в консоль
	fmt.Printf("Файл %s записан", filename)
	//дополнительно в лог
	logger.Printf("Файл %s записан", filename)

	w.WriteHeader(http.StatusOK)

	/*s := fmt.Sprintf("\nMethod: %s\nHost: %s\nPath: %s",
		r.Method, r.Host, r.URL.Path)
	w.Write([]byte(s))
	*/
}
