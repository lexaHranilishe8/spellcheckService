package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spellcheck-service/spellcheck_service"
	"strings"
)

// GetErrors возвращает список ошибок
func GetErrors(w http.ResponseWriter, r *http.Request) {
	errors := spellcheck_service.GetErrors() // Получаем список ошибок

	var formattedErrors []string
	for _, err := range errors {
		words := strings.Fields(err.Name)
		for _, word := range words {
			cleanedWord := strings.Trim(word, `.,!"?`)
			if !spellcheck_service.CheckProducts(cleanedWord) {
				formattedError := fmt.Sprintf("Слово '%s' НЕ найдено в словаре. Идентификатор продукта: %d. Полное имя: %s", cleanedWord, err.ProductID, err.Name)
				formattedErrors = append(formattedErrors, formattedError)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formattedErrors) // Кодируем и отправляем список отформатированных ошибок в формате JSON
}
