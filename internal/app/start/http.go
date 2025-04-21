package start

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"libs/common/grace"
	"libs/common/logger"

	"template/internal/deliveries/course/http/v1"
	"template/internal/services/course"
)

func HTTP(
	courseService course.Service,
	timeoutSeconds int,
	addr string,
	errs chan error,
) grace.Service {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register handlers
	coursev1.Register(e, coursev1.NewV1(courseService, timeoutSeconds))

	// Start server
	go func() {
		logger.Infof(nil, "Starting HTTP server on %s", addr)
		if err := e.Start(addr); err != nil {
			errs <- err
		}
	}()

	return &httpServer{e}
}

type httpServer struct {
	e *echo.Echo
}

func (s *httpServer) Shutdown() error {
	return s.e.Close()
} 