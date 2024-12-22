# go-http-calculator
Yandex Golang practice 
HTTP API калькулятор.

## Возможности

-Базовые арифметические операции(+, -, *, /)
-Поддержка десятичных чисел(3.14)
-Учитывает приорит операций (скобки приоритета, умножение, деление)

## Описание API

### Вычисление выражения

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

## Установка

1. Клонировать репозиторий с помощью `git cline`:
```bash
# ssh
git clone git@github.com:MoodyShoo/go-http-calculator.git
# https 
https://github.com/MoodyShoo/go-http-calculator.git
```
