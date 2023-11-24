package controller

import (
	"os"
	"webhook/pkg/configs"
	"webhook/pkg/worker"

	"github.com/rs/zerolog"
)

type Handler struct {
	config          *configs.Config
	taskDistributor *worker.RedisTaskDistributor
	logger          zerolog.Logger
}

// OpenDBConnection func for opening database connection.
func NewHandler(config *configs.Config, taskDistributor *worker.RedisTaskDistributor) *Handler {

	logger := zerolog.New(os.Stderr).With().Timestamp().Str("Service", "API Handler").Logger()

	return &Handler{config, taskDistributor, logger}
}
