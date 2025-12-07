# TodoApp

Простое приложение для управления задачами.

## Функционал
- Добавление, отметка выполнения и удаление задач
- Синхронизация с сервером

## Стек технологий
- Frontend: React
- Backend: Go (Golang) 
- База данных: MongoDB Atlas

## Запуск

**Backend:**
1. `go mod tidy`
2. Установить `MONGODB_URI` в `.env`
3. `go run ./cmd/server/main.go`

**Frontend:**
1. `cd client`
2. `npm install`
3. `npm start`
