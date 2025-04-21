package course

import (
	"context"
	"libs/apperror"
	"template/internal/repositories/course"
	"template/pkg/domain"
	"template/pkg/reqresp"
)

type CreateCourseRepository interface {
	Create(ctx context.Context, course *domain.Course) error
}

type GetCourseRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Course, error)
}

type ListCoursesRepository interface {
	List(ctx context.Context) ([]*domain.Course, error)
}

type DeleteCourseRepository interface {
	Delete(ctx context.Context, id string) error
}

type SearchCoursesRepository interface {
	Search(ctx context.Context, query string) ([]*domain.Course, error)
}

func CreateCourse(
	ctx context.Context,
	repo CreateCourseRepository,
	req reqresp.CreateCourseRequest,
) (reqresp.CreateCourseResponse, error) {
	course := &domain.Course{
		ID:          generateID(),
		Title:       req.Title,
		Description: req.Description,
		Author:      req.Author,
		Price:       req.Price,
	}

	err := repo.Create(ctx, course)
	if err != nil {
		return reqresp.CreateCourseResponse{}, err
	}

	return reqresp.CreateCourseResponse{
		ID: course.ID,
	}, nil
}

func GetCourse(
	ctx context.Context,
	repo GetCourseRepository,
	id string,
) (reqresp.GetCourseResponse, error) {
	course, err := repo.GetByID(ctx, id)
	if err != nil {
		return reqresp.GetCourseResponse{}, err
	}

	if course == nil {
		return reqresp.GetCourseResponse{}, apperror.NewNotFoundError("course not found")
	}

	return reqresp.GetCourseResponse{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		Author:      course.Author,
		Price:       course.Price,
	}, nil
}

func ListCourses(
	ctx context.Context,
	repo ListCoursesRepository,
) (reqresp.ListCoursesResponse, error) {
	courses, err := repo.List(ctx)
	if err != nil {
		return reqresp.ListCoursesResponse{}, err
	}

	response := reqresp.ListCoursesResponse{
		Courses: make([]reqresp.GetCourseResponse, 0, len(courses)),
	}

	for _, course := range courses {
		response.Courses = append(response.Courses, reqresp.GetCourseResponse{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			Author:      course.Author,
			Price:       course.Price,
		})
	}

	return response, nil
}

func DeleteCourse(
	ctx context.Context,
	repo DeleteCourseRepository,
	id string,
) error {
	return repo.Delete(ctx, id)
}

func SearchCourses(
	ctx context.Context,
	repo SearchCoursesRepository,
	req reqresp.SearchCoursesRequest,
) (reqresp.SearchCoursesResponse, error) {
	courses, err := repo.Search(ctx, req.Query)
	if err != nil {
		return reqresp.SearchCoursesResponse{}, err
	}

	response := reqresp.SearchCoursesResponse{
		Courses: make([]reqresp.GetCourseResponse, 0, len(courses)),
	}

	for _, course := range courses {
		response.Courses = append(response.Courses, reqresp.GetCourseResponse{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			Author:      course.Author,
			Price:       course.Price,
		})
	}

	return response, nil
}

func generateID() string {
	// In a real application, you would use a proper ID generation strategy
	return "course-" + uuid.New().String()
} 