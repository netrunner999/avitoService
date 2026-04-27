package avitoservice

import (
	"context"
	"net/http"
)

// NOTE: интерфейс слоя http-хэндлеров
type Handler interface {
	Test() http.HandlerFunc
}

// NOTE: интерфейс бизнес-логики
type UseCase interface {
	GetTestMessage(ctx context.Context) (string, error)
}

// NOTE: интерфейс базы данных (слой данных)
type Repository interface {
	GetTestMessage(ctx context.Context) (string, error)
}
