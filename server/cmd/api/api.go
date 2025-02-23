package api

import (
	"hexxcore/config"
	"hexxcore/services/attendance"
	"hexxcore/services/auth"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
func (s *APIServer) Run() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Envs.URLS, // Allow all origins (change for security)
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	subrouter := app.Group("/api/v1")
	subrouter.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	userStore := auth.NewStore(s.db)
	authHandler := auth.NewHandler(userStore)
	authHandler.RegisterRoutes(subrouter)

	attendanceStore := attendance.NewStore(s.db)
	attendanceHandler := attendance.NewHandler(attendanceStore)
	attendanceHandler.RegisterRoutes(subrouter)

	log.Fatal(app.Listen(":" + s.addr))
}
