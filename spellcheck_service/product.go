package spellcheck_service

import (
	"encoding/json"
	."errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Product представляет структуру продукта, получаемую из API.
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ProductError представляет структуру ошибки для продуктов.
type ProductError struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
}

// GetProducts запрашивает данные о продуктах через внешний API.
func GetProducts() ([]Product, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", "https://www.sima-land.ru/api/v5/item", nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return nil, err
	}

	// Получаем токен из переменной окружения
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		return nil, errors.New("токен API не установлен. Установите API_TOKEN в переменных окружения")
	}

	// Устанавливаем заголовки запроса
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))
	req.Header.Set("accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Не удалось получить продукты. Статус-код:", resp.StatusCode)
		return nil, errors.New("failed to fetch products")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		return nil, err
	}

	var products []Product
	if err := json.Unmarshal(body, &products); err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return nil, err
	}

	fmt.Printf("Количество полученных продуктов: %d\n", len(products))
	return products, nil
}
