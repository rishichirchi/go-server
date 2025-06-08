package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishichirchi/go-server/database"
	models "github.com/rishichirchi/go-server/model"
)

func ListServers(c *gin.Context){
	var servers []models.MCP

	if err := database.DB.Find(&servers).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve servers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"servers": servers})
}

func GetServer(c *gin.Context){
	var server models.MCP
	name := c.Param("name")

	if err := database.DB.Where("name = ?", name).First(&server).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Server not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"server": server})
}

func AddServer(c *gin.Context){
	var server models.MCP

	if err := c.ShouldBindJSON(&server); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Create(&server).Error; err != nil {
		log.Println("Error adding server:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add server"})
		return
	}


	log.Println("Server added successfully:", server.Name)
	c.JSON(http.StatusCreated, gin.H{"message": "Server added successfully", "server": server})
}

