package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xadichamahkamova/fitness-tracking-app/internal/http/handler"
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
)

func NewGin(queries *storage.Queries) *gin.Engine {

	r := gin.Default()

	handler := handler.HandlerST{
		Queries: queries,
	}
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUser)
	r.GET("/users", handler.ListUsers)
	r.PUT("/users/:id", handler.UpdateUser)
	r.DELETE("/users/:id", handler.DeleteUser)

	return r
}
