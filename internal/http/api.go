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
		
		r.POST("/users/register", handler.Register)
		r.POST("/users/login", handler.LoginUser)
		
		r.GET("/users/:id", handler.GetUser)
		r.GET("/users", handler.ListUsers)
		r.PUT("/users/:id", handler.UpdateUser)
		r.DELETE("/users/:id", handler.DeleteUser)

		r.POST("/password-reset/request", handler.PasswordResetRequest)
		r.POST("/password-reset/verify", handler.VerifyResetToken)
		r.POST("/password-reset/reset", handler.ResetPassword) 

		r.POST("/workouts", handler.CreateWorkout)
		r.GET("/workouts/:id", handler.GetWorkoutByID)
		r.GET("/workouts/user/:id", handler.GetWorkoutByUserID)
		r.PUT("/workouts/:id", handler.UpdateWorkout)
		r.DELETE("/workouts/:id", handler.DeleteWorkout)

		return r
	}
