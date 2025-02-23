package attendance

import (
	"hexxcore/middleware"
	"hexxcore/types"
	"hexxcore/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.AttendanceStore
}

func NewHandler(s types.AttendanceStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	api := router.Group("/attendance")
	api.Use(middleware.AuthMiddleware())

	api.Get("/today", h.GetTodaysClasses)
	api.Get("/subjects", h.GetAllSubjects)
	api.Post("/mark", h.MarkAttendance)
	api.Get("/summary", h.GetSummary)
}
func (h *Handler) GetSummary(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)
	summary, err := h.store.GetAttendanceSummary(userId)
	if err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, fiber.Map{"message": "success", "summary": summary})
}

func (h *Handler) MarkAttendance(c *fiber.Ctx) error {
	req := new(types.AttendanceRequest)
	if err := c.BodyParser(req); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}
	userId := c.Locals("userId").(uint)
	if err := h.store.MarkAttendance(req, userId); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, err)
	}

	return utils.WriteJSON(c, http.StatusOK, fiber.Map{"message": "attendance marked"})
}

func (h *Handler) GetAllSubjects(c *fiber.Ctx) error {
	classes, err := h.store.GetAllSubjects()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
	}
	return utils.WriteJSON(c, http.StatusOK, fiber.Map{"message": "success", "subjects": classes})

}

func (h *Handler) GetTodaysClasses(c *fiber.Ctx) error {
	classes, err := h.store.GetTodaysClasses()
	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, err)
	}
	return utils.WriteJSON(c, http.StatusOK, fiber.Map{"message": "success", "subjects": classes})
}
