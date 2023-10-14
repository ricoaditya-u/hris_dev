package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricoaditya-u/hris_dev/db"
	"github.com/ricoaditya-u/hris_dev/models"
)

func RoleIndex(c *gin.Context) {
	var role []models.Role
	err := db.DB.Find(&role).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": role,
	})
}
