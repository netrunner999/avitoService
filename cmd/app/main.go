package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	httpDelivery "avitoService/internal/avitoService/delivery/http"
	"avitoService/internal/avitoService/repository"
	"avitoService/internal/avitoService/usecase"
)

func main() {
	// создаем слои
	repo := repository.New()
	uc := usecase.New(repo)
	h := httpDelivery.New(uc)

	// регаем маршрут Get /test
	mux := http.NewServeMux()
	mux.HandleFunc("/test", h.Test())

	// конфигурируем http-сервер
	srv := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	// канал для получения сигнала остановки приложения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Server started on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// ждем сигнала остановки
	<-quit
	log.Println("Shutting down the server...")

	// даем 5 секунд на завершение работы
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
	log.Println("Server stopped gracefully.")
}

