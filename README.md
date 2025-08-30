# gRPC Order Book Service

Микросервисное приложение для управления книгами и заказами, построенное на основе gRPC и HTTP API Gateway для переадресации запросов.

##  Архитектура

Приложение состоит из трех основных компонентов:

###  Book Service
- **Технологии**: C# (.NET 8), gRPC, PostgreSQL
- **Функции**: Управление книгами (добавление, получение списка, поиск по ID)

###  Order Service  
- **Технологии**: Go, gRPC, MongoDB
- **Функции**: Управление заказами (создание, получение списка)

###  API Gateway
- **Технологии**: Go, HTTP REST API
- **Функции**: Единая точка входа для клиентов, проксирование запросов к микросервисам

##  Базы данных

- **PostgreSQL**: Хранение данных о книгах
- **MongoDB**: Хранение данных о заказах

##  Быстрый запуск

### Предварительные требования

- Docker и Docker Compose


### Запуск через Docker Compose

1. **Клонируйте репозиторий:**
```bash
git clone github.com/PavelParvadov/grpc_order_book_service
cd grpc_order_book_service
```

2. **Создайте файл .env с переменными окружения:**
```bash
# config
CONFIG_PATH=./config.yaml

# Order service
ORDER_GRPC_PORT=5554
ORDER_GRPC_HOST=order-service

# MongoDB (order-db)
MONGO_DB_PORT=27017
MONGO_DB_HOST=order-db
MONGO_DB_USERNAME=mongo
MONGO_DB_PASSWORD=mongo

# Book service
BOOK_SERVICE_HOST=book-service
BOOK_SERVICE_PORT=5555

# Postgres (book-db)
POSTGRES_DB_PORT=5432
POSTGRES_DB_HOST=postgres
POSTGRES_DB_NAME=book_db
POSTGRES_DB_PASSWORD=postgres
POSTGRES_DB_USERNAME=postgres

```

3. **Запустите все сервисы:**
```bash
docker-compose up --build
```

##  API Endpoints

### Книги

#### Получить список всех книг
```http
GET http://localhost:8080/books
```

**Ответ:**
```json
[
  {
    "id": 1,
    "author": "Джордж Оруэлл",
    "name": "1984"
  }
]
```

#### Добавить новую книгу
```http
POST http://localhost:8080/book
Content-Type: application/json

{
  "author": "Джордж Оруэлл",
  "name": "1984"
}
```

**Ответ:**
```json
{
  "id": 1
}
```

### Заказы

#### Получить список всех заказов
```http
GET http://localhost:8080/orders
```

**Ответ:**
```json
[
  {
    "order_id": "order_123",
    "book_id": 1,
    "status": "PENDING",
    "price": 29.99,
    "place": "Москва"
  }
]
```

#### Создать новый заказ
```http
POST http://localhost:8080/order
Content-Type: application/json

{
  "book_id": 1,
  "status": "PENDING",
  "price": 29.99,
  "place": "Москва"
}
```

**Ответ:**
```json
{
  "order_id": "68766a014bb9d93774596b6c"
}
```


