# OZON API Integration

Интеграция с API OZON для получения и обработки данных о товарах, заказах и остатках.

## Возможности

- Получение списка товаров
- Получение информации о заказах
- Обновление остатков товаров
- Получение аналитических данных
- Логирование всех операций

## Установка

1. Клонируйте репозиторий:

```bash
git clone https://github.com/yourusername/ozon-api-integration.git
cd ozon-api-integration
```

2. Создайте файл окружения:

```bash
touch .env
```

3. Заполните переменные окружения в файле .env:

```bash
OZON_CLIENT_ID=your_client_id
OZON_API_KEY=your_api_key
LOG_LEVEL=info
```

4. Установите зависимости:

```bash
go mod download
```

## Запуск

go run cmd/main.go


## Или запуск с помощью Docker:

```bash
docker build -t ozon-api .
docker run --env-file .env ozon-api
```

## API Endpoints

```bash
GET /products - получить список товаров
GET /orders - получить список заказов
PUT /stocks - обновить остатки
GET /analytics - получить аналитику
```

## Стек технологий

Go 1.24
REST API
Docker
Логирование

### .env.example:
```bash
OZON_CLIENT_ID=your_client_id_here
OZON_API_KEY=your_api_key_here
LOG_LEVEL=info
PORT=8080
```