package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// RequireRole returns a middleware that checks if the user has the required role
func RequireRole(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("role")
		if !exists || userRole != role {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient role"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// RequireRoles returns a middleware that checks if the user has one of the allowed roles
func RequireRoles(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, exists := ctx.Get("role")
		if !exists {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: role not found"})
			ctx.Abort()
			return
		}
		for _, r := range roles {
			if userRole == r {
				ctx.Next()
				return
			}
		}
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient role"})
		ctx.Abort()
	}
}
