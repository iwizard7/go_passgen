package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func generatePassword(length int, symbols string) string {
	rand.Seed(time.Now().UnixNano())

	var password strings.Builder
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + symbols

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(chars))
		password.WriteByte(chars[randomIndex])
	}

	return password.String()
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/generate", func(c *gin.Context) {
		length, _ := strconv.Atoi(c.PostForm("length"))
		symbols := c.PostForm("symbols")

		password := generatePassword(length, symbols)
		c.JSON(http.StatusOK, gin.H{"password": password})
	})

	r.Run(":8080")
}
