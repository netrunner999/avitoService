package main

import (
	httpDelivery "avitoService/internal/avitoService/delivery/http"
	"avitoService/internal/avitoService/repository"
	"avitoService/internal/avitoService/usecase"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	repo := repository.New()
	uc := usecase.New(repo)
	h := httpDelivery.New(uc)

	// NOTE: регаем маршрут Get /test
	mux := http.NewServeMux()
	mux.HandleFunc("/test", h.Test())

	// NOTE: конфигурация http-сервер
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// NOTE: буфферизованный канал
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Server started on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// NOTE: ждем сигнала ОС в главной горутине, если сигнала нет горутина блокируется до его появления
	<-quit
	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
	log.Println("Server stopped gracefully.")
}
