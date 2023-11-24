package controller

import (
	"webhook/app/models"
	"webhook/pkg/response"
	"webhook/pkg/worker"

	"github.com/gofiber/fiber/v2"
)

// @Summary Payload
// @Id payload
// @Tags payload
// @Description send data
// @Accept json
// @Produce      json
// @Param body body map[string]any true "body parameter"
// @Success 200  {object}  models.Data
// @Failure 404 {object} response.Response "bad request: validate your input params"
// @Failure 500 {object} response.Response
// @Router /uplink [post]
func (h *Handler) UplinkHandler(c *fiber.Ctx) error {

	// Create payload instance
	var dataPayload = make(map[string]any)

	// Parse payload in request body
	err := c.BodyParser(&dataPayload)
	if err != nil {
		h.logger.Warn().Str("Error", err.Error()).Msg("failed to parse uplink handler request body")
		return response.CodeMessage(c, fiber.StatusInternalServerError, "Unexpected error occurred! Please contact your administrator")
	}

	// Validate and Decode payload
	data, err := models.ValidateAndDecodePayload(dataPayload, h.config.AttributesMaxLimit, h.config.TraitsMaxLimit)
	if err != nil {
		h.logger.Warn().Str("Error", err.Error()).Msg("Failed to validate and decode payload")
		return response.CodeMessage(c, fiber.StatusBadRequest, err.Error())
	}

	err = h.taskDistributor.DistributeSendPayloadDataTask(c.Context(), &worker.SendPayloadData{&data})
	if err != nil {
		h.logger.Warn().Str("Error", err.Error()).Msg("Task Distribution failed")
		return response.CodeMessage(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Data(c, "data payload sent successfully")

}
