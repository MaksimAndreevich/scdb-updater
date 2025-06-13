package utils

import (
	"strings"
)

// FindInText ищет ключ из мапы в тексте
// Возвращает найденное значение и true, если найдено, или nil и false, если не найдено
func FindInText[T any](text string, searchMap map[string]T) (T, bool) {
	text = strings.ToLower(text)

	// Создаем новую мапу с ключами в нижнем регистре
	lowerSearchMap := make(map[string]T)
	for key, value := range searchMap {
		lowerSearchMap[strings.ToLower(key)] = value
	}

	// Разбиваем текст на слова
	words := strings.Fields(text)

	// Проверяем каждое слово
	for _, word := range words {
		// Очищаем слово от знаков препинания и цифр
		cleanWord := strings.Trim(word, ".,0123456789")
		if cleanWord == "" {
			continue // Пропускаем пустые слова и слова только из цифр
		}

		// Проверяем точное совпадение слова
		if value, ok := lowerSearchMap[cleanWord]; ok {
			return value, true
		}
	}

	return *new(T), false
}
