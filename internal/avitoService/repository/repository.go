package repository

import "context"

// NOTE: реализация интерфейса слоя данных, пустой т. к. в первой лабе не нужны бд
type repo struct {
	// TODO: Реализовать базу данных для работы с PostgreSql
}

// NOTE: функция создания нового экземпляра, возвращает указатель(!)
func New() *repo {
	return &repo{}
}

func (r *repo) GetTestMessage(ctx context.Context) (string, error) {
	// WARNING: возвращаем заглушку
	return "Hello!\n", nil
	// TODO: когда будем делать работу с бд, не забыть про обработку ошибок(!)
}
