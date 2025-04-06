# 0. Базовый образ
FROM golang:1.24

# 1. Установка git (если используешь go-модули с GitHub)
RUN apt-get update && apt-get install -y git

# 2. Рабочая директория внутри контейнера
WORKDIR /app

# 3. Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# 4. Копируем весь проект
COPY . .

# 5. Сборка Go-приложения
RUN go build -o app .

# 6. Запуск
CMD ["./app"]