package connections

import (
	"net/http"
	"time"

	"template/internal/app/config"
)

type Connections struct {
	HTTPClient *http.Client
}

func (c *Connections) Close() {
	// Nothing to close for now
}

func New(cfg *config.Config) (*Connections, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}

	return &Connections{
		HTTPClient: httpClient,
	}, nil
}
