package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    router.GET("/cats/random", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "image": "https://your-cat-image.com/random.jpg",
			// test run
        })
    })

    router.Run(":8080")
}
