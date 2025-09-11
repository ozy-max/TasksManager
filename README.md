# Task Manager API (Go + PostgreSQL)

Учебный проект для изучения **Golang-бэкенда**.  
Сервис позволяет управлять задачами (создание, просмотр, редактирование, удаление) с поддержкой авторизации, категорий, фильтрации и деплоя в Docker.

---

## 🚀 Стек технологий
- **Go** — язык разработки  
- **Gin/Chi** — HTTP-фреймворк  
- **PostgreSQL** — база данных  
- **pgx** — драйвер для PostgreSQL  
- **golang-migrate** — миграции  
- **JWT** — авторизация  
- **Testify** — тестирование  
- **Docker & Docker-compose** — контейнеризация  

---

## 📌 Roadmap

### Этап 1: Основы
- [ ] Init project (`go mod init task-manager-go`)  
- [ ] Hello World  
- [ ] CLI mini-project (todo-список в памяти)  

### Этап 2: HTTP API
- [ ] HTTP server (`/ping`)  
- [ ] Router (chi или gin)  
- [ ] CRUD задач (в памяти)  

### Этап 3: PostgreSQL
- [ ] Подключить PostgreSQL (`pgx`)  
- [ ] Миграции (`golang-migrate`)  
- [ ] CRUD задач (через БД)  

### Этап 4: Архитектура
- [ ] Clean Architecture (`handlers → services → repositories`)  
- [ ] Dependency Injection  

### Этап 5: Авторизация
- [ ] Регистрация / Логин  
- [ ] JWT middleware  
- [ ] Доступ к задачам только владельцу  

### Этап 6: Расширение API
- [ ] Категории задач  
- [ ] Фильтрация / сортировка  
- [ ] Пагинация  

### Этап 7: Тестирование
- [ ] Unit-тесты (`testing`, `testify`)  
- [ ] Моки (`testify/mock`)  
- [ ] Интеграционные тесты (PostgreSQL + testcontainers)  

### Этап 8: Деплой и бонус
- [ ] Dockerize app (`Dockerfile`, `docker-compose.yml`)  
- [ ] Makefile (`make run`, `make test`, `make migrate`)  
- [ ] Bonus: Redis cache  
- [ ] Bonus: WebSocket/gRPC для уведомлений  

---

## ⚙️ Запуск проекта

```bash
# Сборка и запуск
go run main.go

# Тесты
go test ./...

# Запуск через Docker
docker-compose up --build