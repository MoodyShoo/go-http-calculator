# go-http-calculator
Yandex Golang practice 
HTTP API калькулятор.

## Описание API

## Вычисление выражения

**Endpont:** `POST /api/v1/calculate`

**Тело запроса:** `Content-Type: application/json`
Запрос:
```json
{
  "expression": "2+2"
}
```
Ответ (Status 200 OK):
```json
{
  "result":4
}
```

Запрос:
```json
{
  "expression": "2+"
}
```
Ответ (Status 422 Unprocessable Entity):
```json
{
  "error":"Expression is not valid"
}
```

Запрос:
```json
{
  "expression": "2 / 0"
}
```
Ответ (Status 422 Unprocessable Entity):
```json
{
  "error":"Expression is not valid"
}
```
