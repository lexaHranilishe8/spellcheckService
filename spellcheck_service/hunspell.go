package spellcheck_service

import (
	"github.com/akhenakh/hunspellgo"
	"strings"
)

// CheckProducts проверяет правильность названий продуктов.

func CheckProducts(phrase string) bool {

	h := hunspellgo.Hunspell("/home/myuser/GolandProjects/spellcheckService/hunspell/ru_RU.aff", "/home/myuser/GolandProjects/spellcheckService/hunspell/ru_RU.dic")

	// Разбиваем фразу на слова
	words := strings.Fields(phrase)
	for _, word := range words {
		// Очищаем слово от знаков препинания
		cleanedWord := strings.Trim(word, `.,!"?`)
		isCorrect := h.Spell(cleanedWord) // Проверяем правильность слова

		if !isCorrect {
			//fmt.Printf("Слово '%s' НЕ найдено в словаре. это либо ошибка либо корявое ограничение | артикул \n", cleanedWord)
			return false
		}
	}

	return true
}
