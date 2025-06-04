package controller

// import (
// 	"AutoBan/internal/middleware"
// 	"AutoBan/internal/usecase"

// 	"github.com/gin-gonic/gin"
// )

// type UserController struct {
// 	userUseCase usecase.UserUseCase
// }

// func NewUserController() *UserController {
// 	userUseCase := usecase.NewUserUseCase()
// 	return &UserController{userUseCase: userUseCase}
// }

// func UserRoutes(router *gin.Engine) {
// 	c := NewUserController()

// 	userGroup := router.Group("/api/v1/users")
// 	{

// 		protected := userGroup.Use(middleware.AuthMiddleware())
// 		{

// 			protected.GET("/me", c.GetProfile)

// 			protected.PUT("/me", c.UpdateProfile)
// 		}

// 		admin := protected.Use(middleware.RequireAdmin())
// 		{

// 		}
// 	}
// }
