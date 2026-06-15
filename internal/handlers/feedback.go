package handlers

import (
	feedbackdto "cat_service/internal/data/dto/feedback"
	feedbackservice "cat_service/internal/services/feedback"
	"cat_service/internal/validators"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FeedbackHandler struct {
	service *feedbackservice.Service
}

func NewFeedbackHandler(service *feedbackservice.Service) *FeedbackHandler {
	return &FeedbackHandler{service: service}
}

// @Summary Save feedback
// @Description Save feedback for a cat
// @Tags feedback
// @Produce plain
// @Param quality query int true "Quality rating (1-5)"
// @Param cute query int true "Cuteness rating (1-5)"
// @Param message query string true "Additional feedback message"
// @Success 201 {string} string "Feedback saved successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/cats/feedback [get]
func (h *FeedbackHandler) Get(c *fiber.Ctx) error {
	var request feedbackdto.SaveRequest

	if err := c.QueryParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid request query")
	}
	if err := validators.ValidateFeedback(request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid query parameters")
	}

	err := h.service.SaveFeedback(request)
	if err != nil {
		if errors.Is(err, feedbackservice.ErrSaveFeedback) {
			return c.Status(fiber.StatusInternalServerError).SendString("could not save feedback")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
	}
	return c.Status(http.StatusCreated).SendString("feedback saved successfully")
}
