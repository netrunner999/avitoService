package usecase

import (
	domain "avitoService/internal/avitoService/domain"
	"context"
)

// NOTE: реализация интерфейса UseCase через структуру useCase
type useCase struct {
	repo domain.Repository
}

// NOTE: конструктор useCase, возвращает указатель на useCase(!)
func New(repo domain.Repository) *useCase {
	return &useCase{repo: repo}
}

// NOTE: реализация метода интерфейса UseCase
func (u *useCase) GetTestMessage(ctx context.Context) (string, error) {
	return u.repo.GetTestMessage(ctx)
}
