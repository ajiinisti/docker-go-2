package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Ambil nilai environment variable "PORT"
	port := os.Getenv("PORT")
	if port == "" {
		// Jika environment variable "PORT" tidak terdefinisi, gunakan port 8080 sebagai default
		port = "8080"
	}

	// Inisialisasi router Gin
	router := gin.Default()

	// Definisikan route untuk endpoint "/"
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// Jalankan server HTTP menggunakan port yang telah ditentukan
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
