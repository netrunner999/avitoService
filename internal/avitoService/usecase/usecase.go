package usecase

import (
	"context"
	avitoService "avitoService/internal/avitoService"
)

// реализация интерфейса UseCase
type useCase struct {
	repo avitoService.Repository
}

// конструктор useCase, возвращает указатель на useCase(!)
func New(repo avitoService.Repository) *useCase {
	return &useCase{repo: repo}
}

// реализация метода интерфейса UseCase
func (u *useCase) GetTestMessage(ctx context.Context) (string, error) {
	return u.repo.GetTestMessage(ctx)
}
