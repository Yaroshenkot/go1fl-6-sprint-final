package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// автоматическое определение формата и конвертация
func AutoDetectAndConvert(input string) (string, error) {
	input = strings.TrimSpace(input) // удаление пробелов
	if input == "" {
		return "", nil
	}

	//проверка,являтся ли строка кодом М.
	if isMorseCode(input) {
		result := morse.DefaultConverter.ToText(input)
		return result, nil // из М. в текст
	}
	result := morse.DefaultConverter.ToMorse(input)
	return result, nil // из текста в М.
}

func isMorseCode(input string) bool {

	for _, r := range input {
		if !(r == '.' || r == '-' || r == ' ' || r == '\t' || r == '\n' || r == '\r') {
			return false
		}
	}
	// проверка,что строка не состоит только из пробелов
	trimmed := strings.ReplaceAll(input, " ", "")
	trimmed = strings.ReplaceAll(trimmed, "\t", "")
	trimmed = strings.ReplaceAll(trimmed, "\n", "")
	trimmed = strings.ReplaceAll(trimmed, "\r", "")

	if trimmed == "" {
		return false
	}

	hasMorseSymbol := false // проверка, что есть точка или тире
	for _, r := range trimmed {
		if r == '.' || r == '-' {
			hasMorseSymbol = true
			break
		}
	}
	return hasMorseSymbol
}
