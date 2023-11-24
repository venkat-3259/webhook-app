package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

func InitTaskProcessor(redisOpt asynq.RedisClientOpt) {
	taskProcessor := NewRedisTaskProcessor(redisOpt)

	err := taskProcessor.Start()
	if err != nil {
		return
	}
}

type TaskProcessor interface {
	Start() error
	ProcessTaskSendPayload(
		ctx context.Context,
		task *asynq.Task,

	) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
}

func NewRedisTaskProcessor(redisOpt asynq.RedisConnOpt) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{},
	)
	return &RedisTaskProcessor{
		server: server,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(SendData, processor.ProcessTaskSendPayload)

	return nil
}
