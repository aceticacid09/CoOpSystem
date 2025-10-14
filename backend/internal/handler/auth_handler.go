package handler

import (
    "coop/internal/interfaces"
    "coop/internal/models"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your-secret-key") // TODO: Move to environment variable

type AuthHandler struct {
    userService interfaces.UserService
}

func NewAuthHandler(userService interfaces.UserService) *AuthHandler {
    return &AuthHandler{
        userService: userService,
    }
}

// Login authenticates a user
func (h *AuthHandler) Login(c *gin.Context) {
    var loginReq models.LoginRequest

    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    loginResp, err := h.userService.Login(c.Request.Context(), loginReq)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, loginResp)
}

// RefreshToken generates a new access token from a refresh token
func (h *AuthHandler) RefreshToken(c *gin.Context) {
    var req struct {
        RefreshToken string `json:"refresh_token" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // Parse and validate refresh token
    token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
        return
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Check if it's a refresh token
        if tokenType, exists := claims["type"].(string); !exists || tokenType != "refresh" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token type"})
            return
        }

        userIDFloat := claims["user_id"].(float64)
        userID := int(userIDFloat)

        // Get user to get current role
        user, err := h.userService.GetUserByID(c, userID)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            return
        }

        // Create new access token
        newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "user_id":  user.UserID,
            "username": user.Username,
            "role":     user.Role,
            "exp":      time.Now().Add(time.Hour * 24).Unix(), // 24 hours
        })

        newTokenString, err := newToken.SignedString(jwtSecret)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        // Create new refresh token
        newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "user_id": user.UserID,
            "exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
            "type":    "refresh",
        })

        newRefreshTokenString, err := newRefreshToken.SignedString(jwtSecret)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "token":         newTokenString,
            "refresh_token": newRefreshTokenString,
            "user_id":       user.UserID,
            "role":          user.Role,
            "message":       "Token refreshed successfully",
        })
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
    }
}

// Logout (optional - for token blacklisting if implemented)
func (h *AuthHandler) Logout(c *gin.Context) {
    // In a stateless JWT system, logout is typically handled client-side
    // by removing the token. If you want server-side logout, you'd need
    // to implement token blacklisting with Redis or similar.
    c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}