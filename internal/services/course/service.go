package course

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"template/internal/app/store"
	"template/internal/repositories/course"
)

type Service struct {
	st *store.RepositoryStore
	cl *store.ClientStore
}

func New(st *store.RepositoryStore, cl *store.ClientStore) *Service {
	return &Service{
		st: st,
		cl: cl,
	}
}

func (s *Service) CreateCourse(c echo.Context) error {
	var req struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	course := course.Course{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
	}

	if err := s.st.CourseRepository.Create(c.Request().Context(), course); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, course)
}

func (s *Service) GetCourse(c echo.Context) error {
	id := c.Param("id")
	course, err := s.st.CourseRepository.GetByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Course not found")
	}

	return c.JSON(http.StatusOK, course)
}

func (s *Service) ListCourses(c echo.Context) error {
	courses, err := s.st.CourseRepository.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, courses)
}

func (s *Service) DeleteCourse(c echo.Context) error {
	id := c.Param("id")
	if err := s.st.CourseRepository.Delete(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Course not found")
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Service) SearchCourses(c echo.Context) error {
	var req struct {
		Query string `json:"query"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	courses, err := s.st.CourseRepository.Search(c.Request().Context(), req.Query)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, courses)
} 