package middlewares

import (
	"halosuster/src/entities"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, entities.ErrorResponse{
					Status:  false,
					Message: "Unauthorized",
				})
			}
			// Extract the token from the header
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			// Parse the JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Provide the key for verifying the token's signature
				// (replace with your actual key)
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil || !token.Valid {
				c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
				return nil
			}
			// If the token is valid, proceed with the next middleware/handler

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("jwtClaims", claims)

			}
			return next(c)
		}
	}
}
