package main

import (
	"log"
	"os"

	"github.com/makonheimak/task-service/internal/database"
	"github.com/makonheimak/task-service/internal/task/orm"
	"github.com/makonheimak/task-service/internal/task/repository"
	"github.com/makonheimak/task-service/internal/task/service"
	transportgrpc "github.com/makonheimak/task-service/internal/transport/grpc"
)

func main() {
	// Установка переменных окружения по умолчанию для локального запуска
	if os.Getenv("DATABASE_DSN") == "" {
		os.Setenv("DATABASE_DSN", "host=localhost user=postgres password=yourpassword dbname=tasksdb port=5434 sslmode=disable")
	}
	if os.Getenv("USER_SERVICE_ADDR") == "" {
		os.Setenv("USER_SERVICE_ADDR", "localhost:50052")
	}

	// Инициализация БД
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("❌ Could not connect to database: %v", err)
	}

	// Автомиграция
	if err := db.AutoMigrate(&orm.Task{}); err != nil {
		log.Fatalf("❌ Could not migrate database: %v", err)
	}

	repo := repository.NewTaskRepository(db)
	svc := service.NewService(repo)

	// gRPC клиент Users: берем адрес из переменной окружения
	userClient, conn, err := transportgrpc.NewUserClient(os.Getenv("USER_SERVICE_ADDR"))
	if err != nil {
		log.Fatalf("❌ Failed to connect to users: %v", err)
	}
	defer conn.Close()

	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
