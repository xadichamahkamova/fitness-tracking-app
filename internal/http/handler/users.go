package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
	"github.com/xadichamahkamova/hashing/hash"
	t "github.com/xadichamahkamova/fitness-tracking-app/internal/http/token"
)

func (h *HandlerST) Register(c *gin.Context) {

	req := storage.CreateUserParams{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	password, err := hashing.HashPassword(req.PasswordHash.String)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	req.PasswordHash.String = password

	resp, err := h.Queries.CreateUser(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (h *HandlerST) LoginUser(c *gin.Context) {

	req := storage.LoginUserParams{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Queries.LoginUser(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	token := t.GenereteJWTToken(resp.Email.String)
	c.JSON(200, token)
}

func (h *HandlerST) GetUser(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Queries.GetUser(context.Background(), int32(idInt))
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (h *HandlerST) ListUsers(c *gin.Context) {

	resp, err := h.Queries.ListUsers(context.Background())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (h *HandlerST) UpdateUser(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	req := storage.UpdateUserParams{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	req.ID = int32(idInt)
	err = h.Queries.UpdateUser(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User updated"})
}

func (h *HandlerST) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.Queries.DeleteUser(context.Background(), int32(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User deleted"})
}
