package token

import (
	"articlehub-be/models"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Struct untuk klaim token
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken akan menghasilkan token JWT berdasarkan user
func GenerateToken(user models.User) (string, error) {
	// Membaca secret key dari environment variable atau menggunakan default
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "" // Gantilah dengan secret key Anda yang lebih aman
	}

	// Menentukan waktu kedaluwarsa token
	expirationTime := time.Now().Add(24 * time.Hour) // Token akan berlaku selama 24 jam

	// Membuat klaim token
	claims := &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "articlehub", // Bisa diganti dengan nama aplikasi Anda
		},
	}

	// Membuat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken memverifikasi token dan mengembalikan klaim jika valid
func ValidateToken(tokenString string) (*Claims, error) {
	// Membaca secret key dari environment variable atau menggunakan default
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "" // Gantilah dengan secret key Anda yang lebih aman
	}

	// Mem-parsing token dan memverifikasi tanda tangan
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Memastikan bahwa signing method yang digunakan adalah HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	// Jika ada error atau token tidak valid
	if err != nil {
		return nil, err
	}

	// Memeriksa apakah token valid dan mengembalikan klaimnya
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}