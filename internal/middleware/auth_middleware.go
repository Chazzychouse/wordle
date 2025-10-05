package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// AuthMiddleware validates JWT tokens and adds user info to request context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		// Parse and validate token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(getJWTSecret()), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Add user info to request context
		ctx := r.Context()
		ctx = contextWithUserID(ctx, claims.UserID)
		ctx = contextWithUserEmail(ctx, claims.Email)
		ctx = contextWithUserName(ctx, claims.Name)

		// Continue with the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helper functions to add user info to context
func contextWithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, "user_id", userID)
}

func contextWithUserEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, "user_email", email)
}

func contextWithUserName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, "user_name", name)
}

// Helper functions to get user info from context
func GetUserID(ctx context.Context) string {
	if userID, ok := ctx.Value("user_id").(string); ok {
		return userID
	}
	return ""
}

func GetUserEmail(ctx context.Context) string {
	if email, ok := ctx.Value("user_email").(string); ok {
		return email
	}
	return ""
}

func GetUserName(ctx context.Context) string {
	if name, ok := ctx.Value("user_name").(string); ok {
		return name
	}
	return ""
}

func getJWTSecret() string {
	// In production, use environment variable
	// For now, return a default secret
	return "your-secret-key"
}
