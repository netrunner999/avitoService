package usecase

import (
	avitoService "avitoService/internal/avitoService"
	"context"
)

// NOTE: реализация интерфейса UseCase через структуру useCase
type useCase struct {
	repo avitoService.Repository
}

// NOTE: конструктор useCase, возвращает указатель на useCase(!)
func New(repo avitoService.Repository) *useCase {
	return &useCase{repo: repo}
}

// NOTE: реализация метода интерфейса UseCase
func (u *useCase) GetTestMessage(ctx context.Context) (string, error) {
	return u.repo.GetTestMessage(ctx)
}
