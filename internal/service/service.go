package service

import (
	"errors"
	"myproject/pkg/morse"
	"strings"
)

// IsMorseCode проверяет, является ли строка кодом Морзе
func IsMorseCode(s string) bool {
	// Удаляем все допустимые символы (точки, тире и пробелы)
	morseSimbols := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, ".", ""), "-", ""), " ", "")

	// Проверяем, осталась ли неподходящие символы
	if morseSimbols != "" {
		return false // Символы остались
	}

	return true // Использованы только morse символы
}

func TestPrint(textForm string) (string, error) {

	if len(textForm) == 0 {
		return "", errors.New("передана пустая строка для конвертации")

	}

	if IsMorseCode(textForm) {
		convertedText := string(morse.ToText(textForm))
		return convertedText, nil

	} else {
		convertedMorse := string(morse.ToMorse(textForm))
		return convertedMorse, nil

	}
}
