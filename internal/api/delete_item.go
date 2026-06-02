package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (s *Server) DeleteItem(c *fiber.Ctx) error {
	if err := s.Repository.Delete(c.Context(), c.Params("id")); err != nil {
		log.Errorw("s.Repository.Delete", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{})
}
