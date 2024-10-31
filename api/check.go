package api

import (
	"encoding/json"
	"net/http"
	"spellcheck-service/spellcheck_service"
)

// StartCheck запускает проверку орфографии для всех продуктов.
func StartCheck(w http.ResponseWriter, r *http.Request) {
	products, err := spellcheck_service.GetProducts()
	if err != nil {
		http.Error(w, "Error fetching products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var productErrors []spellcheck_service.ProductError
	for _, product := range products {
		if !spellcheck_service.CheckProducts(product.Name) {
			productErrors = append(productErrors, spellcheck_service.ProductError{
				ProductID: product.ID,
				Name:      product.Name,
			})
		}
	}

	spellcheck_service.SaveErrors(productErrors)  // Сохраняем ошибки
	spellcheck_service.UpdateStats(len(products)) // Обновляем статистику

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Spell check completed",
	})
}
