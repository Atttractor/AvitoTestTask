package handlers

import "github.com/gofiber/fiber/v2"


type allBannersQuerryData struct {
	TagId     int `query:"tag_id"`
	FeatureId int `query:"feature_id"`
	Limit     int `query:"limit"`
	Offset    int `query:"offset"`
}

func (h handler) AllBanners(c *fiber.Ctx) error {
	data := allBannersQuerryData{Limit: -1, Offset: -1}

	if err := c.QueryParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	banners, err := h.DB.GetBannersFiltred(data.FeatureId, data.TagId, data.Limit, data.Offset)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if len(banners) == 0 {
		return fiber.NewError(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(&banners)
}