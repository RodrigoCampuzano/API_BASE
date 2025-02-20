package middleware

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "time"
)

func CORSConfig() gin.HandlerFunc {
    config := cors.Config{
        AllowOrigins:     []string{"http://localhost:4200"},  // Cambia según sea necesario
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }

    return cors.New(config)
}
