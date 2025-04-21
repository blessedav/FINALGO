package store

import (
	"template/internal/app/connections"
	"template/internal/repositories/course"
	"template/internal/repositories/course/memory"
)

type RepositoryStore struct {
	CourseRepository course.Repository
}

func NewRepositoryStore(conns *connections.Connections) *RepositoryStore {
	return &RepositoryStore{
		CourseRepository: memory.New(),
	}
}
