# Описание
Интернет-магазин косметики - RESTful API для управления каталогом косметических товаров. Реализованы базовые CRUD-операции с использованием Go и PostgreSQL.
# Технологический стек
- Язык программирования: Go 1.21
- База данных: PostgreSQL 15
- Веб-фреймворк: Chi Router

# Примеры запросов
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Увлажняющий крем",
    "description": "Для сухой кожи лица",
    "price": 1200.50,
    "brand": "L'Oréal",
    "category": "Уход за лицом",
    "stock_quantity": 100
  }'
```

```bash
curl http://localhost:8080/api/products

curl http://localhost:8080/api/products/1
```

```bash
curl -X DELETE http://localhost:8080/api/products/1
```
