package controller

import (
	"net/http"
	"webhook/app/models"
	"webhook/pkg/response"
	"webhook/pkg/utils"

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
	webhookURL := "https://webhook.site/#!/18fe2e34-338c-4bda-881e-acfe7520d482/72ea7726-134f-4d95-b847-2b314b1c7bf1/1"
	go utils.HttpRequest(data, http.MethodPost, webhookURL, "")
	return response.Data(c, "data payload sent successfully")

}
