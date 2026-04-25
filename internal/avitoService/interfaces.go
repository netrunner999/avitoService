package avitoservice

import (
	"context"
	"net/http"
)

// интерфейс слоя http-хэндлеров
type Handler interface {
	Test() http.HandlerFunc
}

// интерфейс бизнес-логики
type UseCase interface {
	GetTestMessage(ctx context.Context) (string, error)
}

// интерфейс базы данных (слой данных)
type Repository interface {
	GetTestMessage(ctx context.Context) (string, error)
}
