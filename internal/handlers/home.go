package handlers

import (
	homeService "cat_service/internal/services/home"

	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	service *homeService.Service
}

func NewHomeHandler(service *homeService.Service) *HomeHandler {
	return &HomeHandler{
		service: service,
	}
}

// @Summary Get random cat image
// @Description Get a random cat image from the external API and render it on the homepage
// @Tags home
// @Produce html
// @Success 200 {string} string "HTML page with random cat image"
// @Failure 500 {string} string "Internal Server Error"
// @Router / [get]
func (h *HomeHandler) Get(c *fiber.Ctx) error {
	image, err := h.service.GetRandomCat()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("error receiving image")
	}

	return c.Render("index", fiber.Map{
		"Title":    "Это твой рандомный кот в эту минуту",
		"CatImage": image,
	})
}
