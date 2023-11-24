package controller

import (
	"os"
	"webhook/pkg/configs"

	"github.com/rs/zerolog"
)

type Handler struct {
	config *configs.Config
	logger zerolog.Logger
}

// OpenDBConnection func for opening database connection.
func NewHandler(config *configs.Config) *Handler {

	logger := zerolog.New(os.Stderr).With().Timestamp().Str("Service", "API Handler").Logger()

	return &Handler{config, logger}
}
