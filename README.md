# go-http-calculator
Yandex Golang practice 
HTTP API калькулятор.

## Возможности

- Базовые арифметические операции(+, -, *, /)
- Поддержка десятичных чисел(3.14)
- Учитывает приорит операций (скобки приоритета, умножение, деление)
- Есть логирование запросов, результатов и ошибок

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

## Установка и настройка

1. Клонировать репозиторий с помощью `git clone`:
```bash
# ssh
git clone git@github.com:MoodyShoo/go-http-calculator.git
# https 
https://github.com/MoodyShoo/go-http-calculator.git
```
2. Перейти в репозиторий:
```bash
cd calc_go
```
3. По умолчанию сервис запускается на порту 8080.
   - Изменить на Windows:
     ```cmd
     set PORT=3000
     ```
     или
     ```powershell
     $env:PORT=3000;
     ```
   - Изменить в Linux:
      ```bash
       PORT=1234
      ```
5. Запустить приложение:
```bash
go run cmd/main.go
```
6. Остановить приложенме:
   Сочетание клавиш `Ctrl + C`

## Пример использования
-Windows
```cmd
curl -X POST http://127.0.0.1:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"2 + 2\"}"
```
-Linux
```bash
curl -X POST http://127.0.0.1:8080/api/v1/calculate \
-H "Content-Type: application/json" \
-d '{"expression": "2 + 2"}'
```

Ожидаемый результат:
```json
{"result":4}
```

# Тестирование
Тесты находятся в директории `/internal/application/application_test.go`
- Запуск тестов
  - Перейти в директорию:
    ```bash
    cd /internal/application
    ```
  - Запустить тесты:
    ```bash
    go test -v . 
    ```
  - Доступные тесты:
    - Valid Sum
    - Invalid Expression
    - Invalid Request Body
    - Valid Double Sum
    - Expression Priority Order
    - Spaces doesn't matter
    - Has letter in expression
    - Uncomplete expression
  - Запуск отдельных тестов:
    ```bash
    go test -v run=ValidSum
    ```
