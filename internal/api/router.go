package api

import (
	"github.com/gin-gonic/gin"
	"github.com/minhtridinh/trid-profile-go/internal/middleware"
	"github.com/minhtridinh/trid-profile-go/internal/service"
)

// SetupRouter initializes the router and register all routes
func SetupRouter(userService service.UserService) *gin.Engine {
	router := gin.Default()

	// Public routes
	public := router.Group("/api/v1")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/register", RegisterHandler(userService))
			auth.POST("/login", LoginHandler(userService))
		}
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware()) // Apply JWT auth middleware
	{
		users := protected.Group("/users")
		{
			users.GET("/me", GetCurrentUserHandler(userService))
			users.PUT("/me", UpdateUserHandler(userService))
		}

		// Admin routes
		admin := protected.Group("/admin")
		admin.Use(middleware.RoleMiddleware("admin")) // Only admin can access
		{
			admin.GET("/users", GetAllUsersHandler(userService))
			admin.GET("/users/:id", GetUserHandler(userService))
			admin.DELETE("/users/:id", DeleteUserHandler(userService))
		}
	}

	return router
}
