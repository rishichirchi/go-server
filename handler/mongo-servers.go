package handler

import (
    "context"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/rishichirchi/go-server/database"
    models "github.com/rishichirchi/go-server/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

const (
    mcpCollection = "mcp-servers-metadata"
    mcpDatabase   = "mcp-servers" // Change to your DB name
)

func ListServersM(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    collection := database.MongoClient.Database(mcpDatabase).Collection(mcpCollection)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve servers"})
        return
    }
    defer cursor.Close(ctx)

    var servers []models.MCP
    if err := cursor.All(ctx, &servers); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse servers"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"servers": servers})
}

func GetServerM(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    name := c.Param("name")
    collection := database.MongoClient.Database(mcpDatabase).Collection(mcpCollection)

    var server models.MCP
    err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&server)
    if err == mongo.ErrNoDocuments {
        c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve server"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"server": server})
}

func AddServerM(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var server models.MCP
    if err := c.ShouldBindJSON(&server); err != nil {
        log.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    collection := database.MongoClient.Database(mcpDatabase).Collection(mcpCollection)
    _, err := collection.InsertOne(ctx, server)
    if err != nil {
        log.Println("Error adding server:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add server"})
        return
    }

    log.Println("Server added successfully:", server.Name)
    c.JSON(http.StatusCreated, gin.H{"message": "Server added successfully", "server": server})
}