package spellcheck_service

import (
	"encoding/json"
	. "errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Product представляет структуру продукта, получаемую из API.
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA2NDg2MDMsImlhdCI6MTczMDA0MzgwMywianRpIjo2MjM4MDU0LCJuYmYiOjE3MzAwNDM4MDN9.RS5UrJbHB9gbTU4P5PiFEZ6NcdXTAqnlHJrFUxbHR-o")
	req.Header.Set("accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Не удалось получить продукты. Статус-код:", resp.StatusCode)
		return nil, New("failed to fetch products")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		return nil, err
	}

	//fmt.Println("Ответ от API:", string(body)) // Отладочное сообщение для проверки ответа

	var products []Product
	if err := json.Unmarshal(body, &products); err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return nil, err
	}

	fmt.Printf("Количество полученных продуктов: %d\n", len(products))
	return products, nil
}
