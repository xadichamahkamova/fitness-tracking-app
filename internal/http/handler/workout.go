package handler

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
)

func (h *HandlerST) CreateWorkout(c *gin.Context) {

	req := storage.CreateWorkoutParams{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Queries.CreateWorkout(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (h *HandlerST) GetWorkoutByUserID(c *gin.Context) {

	req := c.Param("user-id")
	user_id, err := strconv.Atoi(req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Queries.GetWorkoutByUserID(context.Background(), int32(user_id))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (h *HandlerST) GetWorkoutByID(c *gin.Context) {

	req := c.Param("id")
	workout_id, err := strconv.Atoi(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Queries.GetWorkoutByID(context.Background(), int32(workout_id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}
