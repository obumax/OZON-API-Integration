package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"ozon-api-integration/pkg/logger"
)

// OzonClient представляет собой клиент для взаимодействия с OZON API
type OzonClient struct {
	clientID string
	apiKey   string
	baseURL  string
	client   *http.Client
}

// NewOzonClient создает новый экземпляр OzonClient
func NewOzonClient(clientID, apiKey string) *OzonClient {
	return &OzonClient{
		clientID: clientID,
		apiKey:   apiKey,
		baseURL:  "https://api-seller.ozon.ru",
		client:   &http.Client{},
	}
}

// makeRequest выполняет HTTP-запрос к OZON API
func (c *OzonClient) makeRequest(method, endpoint string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, c.baseURL+endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Client-Id", c.clientID)
	req.Header.Set("Api-Key", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	logger.Info(fmt.Sprintf("Making %s request to %s", method, endpoint))

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}

// GetProducts получает список продуктов из OZON API
func (c *OzonClient) GetProducts() ([]byte, error) {
	return c.makeRequest("POST", "/v2/product/list", map[string]interface{}{
		"filter": map[string]interface{}{
			"visibility": "ALL",
		},
		"limit": 100,
	})
}

// GetOrders получает список заказов из OZON API
func (c *OzonClient) GetOrders() ([]byte, error) {
	return c.makeRequest("POST", "/v3/posting/fbs/list", map[string]interface{}{
		"filter": map[string]interface{}{
			"since": "2024-01-01T00:00:00.000Z",
			"to":    "2024-12-31T23:59:59.999Z",
		},
		"limit": 100,
	})
}

// UpdateStocks обновляет запасы продуктов в OZON API
func (c *OzonClient) UpdateStocks(stocks interface{}) ([]byte, error) {
	return c.makeRequest("POST", "/v1/product/import/stocks", stocks)
}

// GetAnalytics получает аналитические данные из OZON API
func (c *OzonClient) GetAnalytics() ([]byte, error) {
	return c.makeRequest("POST", "/v1/analytics/data", map[string]interface{}{
		"date_from": "2024-01-01",
		"date_to":   "2024-12-31",
		"metrics":   []string{"revenue", "orders_count"},
	})
}
