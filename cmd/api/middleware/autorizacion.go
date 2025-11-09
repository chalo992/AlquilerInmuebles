package middleware

import (
	"AlquilerInmuebles/cmd/api/common"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Autorizacion(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header requerido"})
		c.Abort()
		return
	}

	// Espera formato: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
		c.Abort()
		return
	}

	tokenString := parts[1]

	// Validar el token con la clave secreta
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return []byte(common.Secret), nil
	})

	// Si hay error o el token no es válido
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	// Extraer los claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Guardamos el ID como uint
		if idFloat, ok := claims["id"].(float64); ok {
			c.Set("id", uint(idFloat))
		}
		// Guardamos el rol tal como viene (string)
		if rol, ok := claims["rol"].(string); ok {
			c.Set("rol", rol)
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudieron leer los claims"})
		c.Abort()
		return
	}

	c.Next()
}
