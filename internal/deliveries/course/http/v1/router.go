package v1

import "github.com/labstack/echo/v4"

func Register(e *echo.Echo, h *Handler) {
	e.POST("/v1/courses", h.CreateCourse)
	e.GET("/v1/courses", h.ListCourses)
	e.GET("/v1/courses/:id", h.GetCourse)
	e.DELETE("/v1/courses/:id", h.DeleteCourse)
	e.POST("/v1/courses/search", h.SearchCourses)
} 