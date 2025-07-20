package service

import (
	"fmt"
	"log"
	"myproject/pkg/morse"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode"
)

// IsMorseCode проверяет, является ли строка кодом Морзе
func IsMorseCode(s string) bool {
	// Удаляем пробелы в начале и конце строки
	s = strings.TrimSpace(s)

	// Проверяем каждый символ в строке
	for _, char := range s {
		if !(char == '.' || char == '-' || unicode.IsSpace(char)) {
			return false // Если найден недопустимый символ, возвращаем false
		}
	}
	return true // Если все символы допустимы, возвращаем true
}

func newFilenameTime(oldName string) string {

	name := filepath.Base(oldName)
	ext := filepath.Ext(oldName)
	time := time.Now().UTC().Format("20060102_150405")

	return fmt.Sprint(name[:len(name)-len(ext)], "_", time, ext)

}

func createFolder(dir string) {

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateFile(filename string, text string, dir string) {

	filename = newFilenameTime(filename)

	// формируем относительное имя нужного файла
	if len(dir) > 0 {
		createFolder(dir)
		filename = filepath.Join(dir, filename)
		fmt.Println(filename)
	}

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
	fmt.Println(text)

	// возвращаем обратно вывод в консоль
	os.Stdout = stdout
	// строка выведется в консоль
	fmt.Printf("Файл %s записан", filename)

}

func TestPrint(textForm string) (string, bool) {

	if IsMorseCode(textForm) {
		convertedText := string(morse.ToText(textForm))
		return convertedText, true

	} else {
		convertedMorse := string(morse.ToMorse(textForm))
		return convertedMorse, false

	}
}
