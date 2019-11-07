package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"iiujapp.tech/basic-gin/model"
	"iiujapp.tech/basic-gin/service"
)

// UserHandler function
func UserHandler(s service.Iservice, c *gin.Context) {
	res, err := s.QueryUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{ 
		"code" : http.StatusOK, 
		"message": res,
	})
}

// UserSaveHandler function
func UserSaveHandler(s service.Iservice, c *gin.Context) {
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.WriteData(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{ 
		"code" : http.StatusOK, 
		"message": true,
	})
}
