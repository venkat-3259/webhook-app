package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webhook/app/models"
	"webhook/pkg/utils"

	"github.com/hibiken/asynq"
)

const SendData = "task:send_payload_data"

type SendPayloadData struct {
	*models.Payload
}

func (distributor *RedisTaskDistributor) DistributeSendPayloadDataTask(ctx context.Context, payload *SendPayloadData, opts ...asynq.Option) error {

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Worker: Failed to marshal payload, Reason: %v", err)
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(SendData, jsonPayload, opts...)

	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		log.Printf("Worker: Failed to enqueue task, Reason: %v", err)
		return fmt.Errorf("failed to enqueue task %w", err)
	}

	log.Printf("payload %v, queue %v max_retry %v task enqueued", info.Queue, info.MaxRetry, task.Payload())
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendPayload(ctx context.Context, task *asynq.Task) error {

	var payload SendPayloadData
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		log.Println("Failed to unmarshal payload data, Reason: ", err)
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	webhookURL := "https://webhook.site/ 	//https://webhook.site/18fe2e34-338c-4bda-881e-acfe7520d482"

	statusCode, respBody, err := utils.HttpRequest(payload, http.MethodPost, webhookURL, "")
	if err != nil {
		log.Println("Error while sending http request, Error: ", err)
		return fmt.Errorf("failed to send data %w", err)
	} else {
		fmt.Println("Status Code:", statusCode, "Response:", respBody)
	}
	return nil
}
