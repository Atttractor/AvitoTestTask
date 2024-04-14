package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type createBannerQuerryData struct {
	TagsIds 	[]uint 					`json:"tags_ids"`
	FeatureId 	uint 					`json:"feature_id"`
	Content 	map[string]interface{} 	`json:"content"`
	IsActive 	bool 					`json:"is_active"`
}

func (h handler) CreateBanner(c *fiber.Ctx) error {
	data := createBannerQuerryData{}

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	content, err := contentToString(data.Content)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	id, err := h.DB.CreateBanner(data.TagsIds, data.FeatureId, content, data.IsActive)

	if err != nil || id == -1 {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&id)
}

func contentToString(m interface{}) (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(b), nil
}