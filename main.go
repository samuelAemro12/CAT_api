package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

const catAPIURL = "https://api.thecatapi.com/v1/images/search"

func getCatImage() (string, error) {
    resp, err := http.Get(catAPIURL)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)

    var result []map[string]interface{}
    json.Unmarshal(body, &result)

    return result[0]["url"].(string), nil
}

func main() {
    router := gin.Default()

    router.GET("/cats/random", func(c *gin.Context) {
        imageURL, err := getCatImage()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch image"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"image": imageURL})
    })

    router.Run(":8080")
}
