package handlers

import (
	"encoding/json"
	"net/http"

	"ozon-api-integration/internal/client"
	"ozon-api-integration/internal/config"
	"ozon-api-integration/pkg/logger"
)

// ProductHandler обрабатывает HTTP-запросы, связанные с продуктами и заказами
type ProductHandler struct {
	ozonClient *client.OzonClient
}

// NewProductHandler создает новый экземпляр ProductHandler
func NewProductHandler(cfg *config.Config) *ProductHandler {
	return &ProductHandler{
		ozonClient: client.NewOzonClient(cfg.OzonClientID, cfg.OzonAPIKey),
	}
}

// GetProducts обрабатывает запрос на получение списка продуктов
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	logger.Info("Getting products from OZON API")

	data, err := h.ozonClient.GetProducts()
	if err != nil {
		logger.Error("Failed to get products: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// GetOrders обрабатывает запрос на получение списка заказов
func (h *ProductHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	logger.Info("Getting orders from OZON API")

	data, err := h.ozonClient.GetOrders()
	if err != nil {
		logger.Error("Failed to get orders: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// UpdateStocks обрабатывает запрос на обновление запасов продуктов
func (h *ProductHandler) UpdateStocks(w http.ResponseWriter, r *http.Request) {
	logger.Info("Updating stocks in OZON API")

	var stocks interface{}
	if err := json.NewDecoder(r.Body).Decode(&stocks); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	data, err := h.ozonClient.UpdateStocks(stocks)
	if err != nil {
		logger.Error("Failed to update stocks: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// GetAnalytics обрабатывает запрос на получение аналитических данных
func (h *ProductHandler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	logger.Info("Getting analytics from OZON API")

	data, err := h.ozonClient.GetAnalytics()
	if err != nil {
		logger.Error("Failed to get analytics: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
