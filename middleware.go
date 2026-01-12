package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

// reCAPTCHA middleware
func RecaptchaMiddleware(verify func(token string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("g-recaptcha-response")

		if !verify(token) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Robot detected",
			})
			return
		}

		c.Next()
	}
}
