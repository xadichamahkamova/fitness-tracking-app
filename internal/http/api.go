package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xadichamahkamova/fitness-tracking-app/internal/email"
	"github.com/xadichamahkamova/fitness-tracking-app/internal/http/handler"
	"github.com/xadichamahkamova/fitness-tracking-app/internal/http/middleware"
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
)

func NewGin(queries *storage.Queries, notif email.NotificationRepo) *gin.Engine {

	r := gin.Default()

	handler := handler.HandlerST{
		Queries: queries,
		Notification: notif,
	}
	r.Use(middleware.AuthMiddleware())
	
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUser)
	r.GET("/users", handler.ListUsers)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	r.POST("/password-reset/request", handler.PasswordResetRequest)
	r.POST("/password-reset/verify", handler.VerifyResetToken)
	r.POST("/password-reset/reset", handler.ResetPassword) 

	r.POST("/workout", handler.CreateWorkout)
	r.GET("/workout/:id", handler.GetWorkoutByID)
	r.GET("/workout/:user-id", handler.GetWorkoutByUserID)

	return r
}
