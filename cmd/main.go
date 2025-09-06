package main

import (
	"log"
	"net/http"
	"os"

	"ozon-api-integration/internal/config"
	"ozon-api-integration/internal/handlers"
	"ozon-api-integration/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Инициализация конфигурации
	cfg := config.New()

	// Инициализация логгера
	logger.Init(cfg.LogLevel)

	// Создание обработчиков
	productHandler := handlers.NewProductHandler(cfg)

	// Настройка роутов
	r := mux.NewRouter()
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/orders", productHandler.GetOrders).Methods("GET")
	r.HandleFunc("/stocks", productHandler.UpdateStocks).Methods("PUT")
	r.HandleFunc("/analytics", productHandler.GetAnalytics).Methods("GET")

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Server starting on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
