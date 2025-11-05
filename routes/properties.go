package routes

import (
	"net/http"

	"github.com/Ademayowa/learn-terraform/models"
	"github.com/gin-gonic/gin"
)

func CreateProperty(c *gin.Context) {
	var property models.Property

	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse property data"})
		return
	}

	if err := property.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save property"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "property created", "property": property})
}

func GetProperties(c *gin.Context) {
	properties, err := models.GetAllProperties()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch properties"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"properties": properties})
}
