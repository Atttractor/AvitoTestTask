package handlers

import (
	"dooplik/avitoTestTask/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type getBannerQuerryData struct {
	TagId 			uint 	`query:"tag_id,required"`
	FeatureId 		uint 	`query:"feature_id,required"`
	UseLastVersion 	bool 	`query:"use_last_version"`
}

type getUserBannerResponse struct {
	Id 			uint 	`json:"id"`
	IsActive 	bool 	`json:"is_active"`
	Data 		string 	`json:"data"`
	FeatureId 	uint 	`json:"feature_id"`
	TagsId 		[]uint `json:"tags_ids"`
}

func (h *handler) GetBanner(c *fiber.Ctx) error {
	data := getBannerQuerryData{}

	if err := c.QueryParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	banner := models.Banner{}
	if data.UseLastVersion {
		b, err := h.DB.GetBanner(data.FeatureId, data.TagId)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		banner = b
	} else {
		b, err := h.local.GetBanner(data.TagId, data.FeatureId)
		if err != nil {
			fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		
		banner = b
	}

	if banner.ID <= 0 {
		return fiber.NewError(fiber.StatusNotFound)
	}

	response := getUserBannerResponse{
		Id: banner.ID,
		IsActive: banner.IsActive,
		Data: banner.Data,
		FeatureId: banner.FeatureId,
		TagsId: []uint{},
	}

	for _, tag := range(banner.Tags) {
		response.TagsId = append(response.TagsId, tag.ID)
	}

	return c.Status(fiber.StatusOK).JSON(&response)
}