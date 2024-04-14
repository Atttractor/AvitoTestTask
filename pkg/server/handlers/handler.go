package handlers

import (
	"dooplik/avitoTestTask/pkg/common/storage"
	"dooplik/avitoTestTask/pkg/server/middleware"

	"github.com/gofiber/fiber/v2"
)
type handler struct {
	DB db.Storage
	local db.LocalDB
}

func RegisterRoutes(app *fiber.App, db db.Storage, l db.LocalDB) {
    h := &handler{
        DB: db,
		local: l,
	}

    userGroup := app.Group("/user_banner", middleware.IsUser)
    userGroup.Get("/", h.GetBanner)

    adminGroup := app.Group("/banner", middleware.IsAdmin)
    adminGroup.Get("/", h.AllBanners)
    adminGroup.Post("/", h.CreateBanner)
    //adminGroup.Patch("/:id", mockHandler.Handle)
    //adminGroup.Delete("/:id", mockHandler.Handle)

}