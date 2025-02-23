package auth

import (
	"fmt"
	"hexxcore/config"
	"hexxcore/types"
	"hexxcore/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(s types.UserStore) *Handler {
	return &Handler{store: s}
}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	router := r.Group("/auth")
	router.Post("/signup", h.Signup) // User Signup
	router.Post("/login", h.Login)   // User Login
}

func (h *Handler) Signup(c *fiber.Ctx) error {
	user := new(types.User)
	if err := c.BodyParser(user); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
	}
	if err := config.Validator.Struct(user); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
	}
	// Hash password before storing
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "failed to process the password"})
	}

	user.Password = hash
	user.Role = "student"

	// Email Verification Setup

	// Store user
	if err := h.store.CreateUser(user); err != nil {
		return utils.WriteError(c, http.StatusInternalServerError, fmt.Errorf("error creating user %v", err))
	}

	// Generate JWT token for user
	token := utils.GenerateJWT(user.ID, user.Role)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   c.Protocol() == "https",
		SameSite: "strict",
		HTTPOnly: true,
	})
	user.Password = ""
	return utils.WriteJSON(c, http.StatusOK, map[string]any{"message": "verification email sent successfully", "user": user})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	req := new(types.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	if err := config.Validator.Struct(req); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	user, err := h.store.GetUserByIdentifier(req.Identifier)
	if err != nil {
		return utils.WriteError(c, http.StatusBadRequest, err)
	}

	// validate the password
	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return utils.WriteError(c, http.StatusBadRequest, fmt.Errorf("invalid credentials"))
	}
	// Generate JWT token for user
	token := utils.GenerateJWT(user.ID, user.Role)
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Secure:   c.Protocol() == "https",
		SameSite: "strict",
		HTTPOnly: true,
	})
	user.Password = ""
	return utils.WriteJSON(c, http.StatusOK, fiber.Map{"message": "login successfull", "user": user})
}
