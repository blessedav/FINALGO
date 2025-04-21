package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"libs/common/ctxconst"

	"template/internal/services/course"
	"template/pkg/deliveries"
	"template/pkg/reqresp"
)

type Handler struct {
	service course.Service
	timeout time.Duration
}

func NewV1(service course.Service, timeoutSeconds int) *Handler {
	if timeoutSeconds <= 0 {
		timeoutSeconds = 60
	}

	return &Handler{
		service: service,
		timeout: time.Second * time.Duration(timeoutSeconds),
	}
}

func (h *Handler) CreateCourse(c echo.Context) error {
	var request reqresp.CreateCourseRequest
	if err := c.Bind(&request); err != nil {
		return deliveries.HandleEcho(c, err)
	}

	ctx, cancel := h.context(c)
	defer cancel()

	response, err := h.service.CreateCourse(ctx, request)
	if err != nil {
		return deliveries.HandleEcho(c, err)
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *Handler) GetCourse(c echo.Context) error {
	id := c.Param("id")

	ctx, cancel := h.context(c)
	defer cancel()

	response, err := h.service.GetCourse(ctx, id)
	if err != nil {
		return deliveries.HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) ListCourses(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	response, err := h.service.ListCourses(ctx)
	if err != nil {
		return deliveries.HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteCourse(c echo.Context) error {
	id := c.Param("id")

	ctx, cancel := h.context(c)
	defer cancel()

	err := h.service.DeleteCourse(ctx, id)
	if err != nil {
		return deliveries.HandleEcho(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) SearchCourses(c echo.Context) error {
	var request reqresp.SearchCoursesRequest
	if err := c.Bind(&request); err != nil {
		return deliveries.HandleEcho(c, err)
	}

	ctx, cancel := h.context(c)
	defer cancel()

	response, err := h.service.SearchCourses(ctx, request)
	if err != nil {
		return deliveries.HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) context(c echo.Context) (context.Context, context.CancelFunc) {
	ctx := c.Request().Context()
	ctx = ctxconst.SetRequestID(ctx, "test-request-id")
	ctx = ctxconst.SetUserID(ctx, "test-user-id")
	ctx = ctxconst.SetUserPhoneNumber(ctx, "test-phone-number")

	return context.WithTimeout(ctx, h.timeout)
} 