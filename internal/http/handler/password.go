package handler

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/xadichamahkamova/fitness-tracking-app/internal/http/token"
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
	hashing "github.com/xadichamahkamova/hashing/hash"
)

type ResetPassword struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func (h *HandlerST) PasswordResetRequest(c *gin.Context) {

	req := storage.SavePasswordResetTokenParams{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	req.UserToken = token.GenereteJWTToken(req.UserEmail).RefreshToken

	err := h.Queries.SavePasswordResetToken(context.Background(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = h.Notification.SendEmail(req.UserEmail, req.UserToken)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Password reset email sent"})
}

func (h *HandlerST) VerifyResetToken(c *gin.Context) {

	var req string
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := token.ExtractClaim(req)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Token verified, proceed with password reset"})
}

func (h *HandlerST) ResetPassword(c *gin.Context) {

	req := ResetPassword{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	claims, err := token.ExtractClaim(req.Token)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	email := claims["user_email"].(string)

	hashedPassword, err := hashing.HashPassword(req.NewPassword)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = h.Queries.UpdateUserPassword(context.Background(), storage.UpdateUserPasswordParams{
		Email:        sql.NullString{String: email, Valid: true},
		PasswordHash: sql.NullString{String: hashedPassword, Valid: true},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Password updated successfully"})
}
