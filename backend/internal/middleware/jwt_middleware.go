package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Context key to store claims
type contextKey string

const ClaimsContextKey contextKey = "claims"

// JWTMiddleware checks if the request has a valid JWT token
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header
		authHeader := r.Header.Get("Authorization")

		// Check if Authorization Header exists and starts with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, `{"error": "missing_token"}`, http.StatusUnauthorized)
			return
		}

		// Extract token (removing "Bearer " prefix)
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate Token
		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusUnauthorized)
			return
		}

		// Store claims in request context and pass to the next handler
		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ValidateToken parses and validates a JWT token
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Get JWT Secret Key from environment variables
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, fmt.Errorf("server_error: missing_JWT_SECRET")
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the token's signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid_token")
		}
		return []byte(secretKey), nil
	})

	// Check if token is valid
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid_token")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid_token")
	}

	// Check expiration
	expiration, ok := claims["exp"].(float64) // JWT exp is stored as float64
	if !ok {
		return nil, fmt.Errorf("invalid_token")
	}

	if time.Now().Unix() > int64(expiration) {
		return nil, fmt.Errorf("token_expired")
	}

	return claims, nil
}
