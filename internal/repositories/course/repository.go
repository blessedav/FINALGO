package course

import (
	"context"

	"template/pkg/domain"
)

type Repository interface {
	Create(ctx context.Context, course *domain.Course) error
	GetByID(ctx context.Context, id string) (*domain.Course, error)
	List(ctx context.Context) ([]*domain.Course, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string) ([]*domain.Course, error)
}

func NewMemoryRepository() Repository {
	return memory.New()
} 