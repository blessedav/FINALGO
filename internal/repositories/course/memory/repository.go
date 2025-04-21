package memory

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"template/pkg/domain"
)

type repository struct {
	mu      sync.RWMutex
	courses map[string]*domain.Course
}

func New() *repository {
	return &repository{
		courses: make(map[string]*domain.Course),
	}
}

func (r *repository) Create(ctx context.Context, course *domain.Course) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	course.ID = uuid.New().String()
	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()

	r.courses[course.ID] = course
	return nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*domain.Course, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	course, exists := r.courses[id]
	if !exists {
		return nil, nil
	}
	return course, nil
}

func (r *repository) List(ctx context.Context) ([]*domain.Course, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	courses := make([]*domain.Course, 0, len(r.courses))
	for _, course := range r.courses {
		courses = append(courses, course)
	}
	return courses, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.courses, id)
	return nil
}

func (r *repository) Search(ctx context.Context, query string) ([]*domain.Course, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var results []*domain.Course
	for _, course := range r.courses {
		if contains(course.Title, query) || contains(course.Description, query) {
			results = append(results, course)
		}
	}
	return results, nil
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
} 