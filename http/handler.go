package http

import "github.com/gofiber/fiber/v2"

type Handler struct {
}

func (h Handler) Register(router fiber.Router) {
	router.Post("", h.registerEvents)
	router.Get("/:eventId", h.getEvents)
	router.Get("/:eventId/status", h.getEventStatus)
	router.Get("/url", h.registerUrl)
}

func (h Handler) registerEvents(ctx *fiber.Ctx) error {
	return ctx.SendString("register events")
}

func (h Handler) getEvents(ctx *fiber.Ctx) error {
	return ctx.SendString("get events")
}

func (h Handler) getEventStatus(ctx *fiber.Ctx) error {
	return ctx.SendString("get event status")
}

func (h Handler) registerUrl(ctx *fiber.Ctx) error {
	return ctx.SendString("register url")
}
