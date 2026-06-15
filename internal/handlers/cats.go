package handlers

import (
	catdto "cat_service/internal/data/dto/cats"
	catservice "cat_service/internal/services/cats"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CatHandler struct {
	service *catservice.Service
}

func NewCatHandler(service *catservice.Service) *CatHandler {
	return &CatHandler{service: service}
}

// @Summary Get cats
// @Description Get list of all cats preview with optional limit
// @Tags cats
// @Produce json
// @Param limit query int false "Limit the number of cats returned"
// @Success 200 {array} catdto.CatPreview
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/cats [get]
func (h *CatHandler) Get(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 0)

	cats, err := h.service.GetCats(limit)
	if err != nil {
		if errors.Is(err, catservice.ErrGetCats) {
			return c.Status(fiber.StatusInternalServerError).SendString("failed to get cats")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
	}
	return c.Status(fiber.StatusOK).JSON(cats)
}

// @Summary Get cat by ID
// @Description Get detailed information about a cat by its ID
// @Tags cats
// @Produce json
// @Param id path int true "Cat ID"
// @Success 200 {object} catdto.CatResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/cats/{id} [get]
func (h *CatHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid input parameters")
	}

	cat, err := h.service.GetCat(id)
	if err != nil {
		if errors.Is(err, catservice.ErrGetCat) {
			return c.Status(fiber.StatusInternalServerError).SendString("failed to get cat")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
	}
	return c.Status(fiber.StatusOK).JSON(cat)
}

// @Summary Get cat view by ID
// @Description Get a rendered view of a cat by its ID
// @Tags cats
// @Produce html
// @Param id path int true "Cat ID"
// @Success 200 {string} string "HTML view of the cat"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /views/cats/internal/{id} [get]
func (h *CatHandler) GetViewByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid input parameters")
	}

	cat, err := h.service.GetCat(id)
	if err != nil {
		if errors.Is(err, catservice.ErrGetCat) {
			return c.Status(fiber.StatusInternalServerError).SendString("failed to get cat")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
	}

	homeless := "нет"
	if cat.Homeless {
		homeless = "да"
	}

	return c.Render("cat_prev", fiber.Map{
		"ID":        cat.ID,
		"Name":      cat.Name,
		"Age":       cat.Age,
		"CatImage":  cat.ImageUrl,
		"Homeless":  homeless,
		"CreatedAt": cat.CreatedAt,
	})
}

// @Summary Create a new cat
// @Description Create a new cat with the provided information
// @Tags cats
// @Accept json
// @Produce json
// @Param cat body catdto.CatRequest true "Cat information"
// @Success 201 {object} catdto.CatResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/cats [post]
func (h *CatHandler) Post(c *fiber.Ctx) error {
	var payload catdto.CatRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid parse body")
	}

	cat, err := h.service.SaveCat(payload)
	if err != nil {
		if errors.Is(err, catservice.ErrSaveCat) {
			return c.Status(fiber.StatusInternalServerError).SendString("could not save cat")
		}
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
	}
	return c.Status(http.StatusCreated).JSON(cat)
}
