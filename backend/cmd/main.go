package main

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    
    // ヘルスチェック用エンドポイント

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "OK",
            "message": "Go backend is running!",
        })
    })
    

    // APIルート例
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to Third Opinion API",
            "version": "1.0.0",
        })
    })
    
    r.Run(":8080")
}
