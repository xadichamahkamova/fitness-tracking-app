package handler

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xadichamahkamova/fitness-tracking-app/internal/http/token"
	"github.com/xadichamahkamova/fitness-tracking-app/storage"
	hashing "github.com/xadichamahkamova/hashing/hash"
)

func (h *HandlerST) PasswordResetRequest(c *gin.Context) {

	type request struct {
		Email string `json:"email"`
	}

	var input request
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resetToken := token.GenereteJWTToken(input.Email).RefreshToken

	err := h.Queries.SavePasswordResetToken(context.Background(), storage.SavePasswordResetTokenParams{
		UserEmail: input.Email,
		UserToken: resetToken,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.Notification.SendEmail(input.Email, resetToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent"})
}

func (h *HandlerST) VerifyResetToken(c *gin.Context) {

	type request struct {
		Token string `json:"token"`
	}

	var input request
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := token.ExtractClaim(input.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token verified, proceed with password reset"})
}

func (h *HandlerST) ResetPassword(c *gin.Context) {

	type request struct {
		Token       string `json:"token"`
		NewPassword string `json:"new_password"`
	}

	var input request
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := token.ExtractClaim(input.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	email := claims["user_email"].(string)

	hashedPassword, err := hashing.HashPassword(input.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.Queries.UpdateUserPassword(context.Background(), storage.UpdateUserPasswordParams{
		Email:        sql.NullString{String: email, Valid: true},
		PasswordHash: sql.NullString{String: hashedPassword, Valid: true},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
