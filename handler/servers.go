package handler

import (
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