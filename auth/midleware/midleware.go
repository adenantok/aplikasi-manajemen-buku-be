package midleware

import (
	"aplikasi-manajemen-buku-be/auth/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenuser := c.GetHeader("Authorization")
		if tokenuser == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Validasi token di sini
		claims, err := token.ValidateToken(tokenuser) // Menggunakan error untuk menentukan validitas token
		if err != nil {                               // Jika ada error, token invalid
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Simpan userID di context untuk digunakan di controller
		c.Set("userID", claims.UserID)
		c.Next()
	}
}

// CORS Middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		// Handle preflight request
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
