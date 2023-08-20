package controllers

import (
	"net/http"
	// "log"
	"github.com/Akhilstaar/me-my_encryption/db"
	"github.com/Akhilstaar/me-my_encryption/models"
	"github.com/Akhilstaar/me-my_encryption/utils"
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
)

var Db db.PuppyDb

func AddNewUser(c *gin.Context) {
	// TODO: Modify this function to handle multiple concatenated json inputs


	// TODO: Implement admin authentication logic
	// Authenticate the admin here

	// Validate the input format
	info := new(models.TypeUserNew)
	if err := c.BindJSON(info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input data format."})
		return
	}

	// Create user
	newUser := models.User{
		Id:      info.Id,
		Name:    info.Name,
		Email:   info.Email,
		Gender:  info.Gender,
		Pass:    info.PassHash,
		PrivK:   "",
		PubK:    "",
		AuthC:   utils.RandStringRunes(15),
		Data:    "",
		Submit:  false,
		Matches: "",
		Dirty:   true,
	}

	// Insert the user into the database

	if err := Db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

func DeleteUser(c *gin.Context) {
	// TODO: Implement admin authentication logic
	// Authenticate the admin here

	// Validate the input format
	info := new(models.TypeUserNew)
	if err := c.BindJSON(info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input data format."})
		return
	}

	newUser := models.User{
		Id:     info.Id,
		Name:   info.Name,
		Email:  info.Email,
		Gender: info.Gender,
	}

	if err := Db.Unscoped().Delete(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User Deleted successfully."})
}

func DeleteAllUsers(c *gin.Context) {
	// TODO: Implement admin authentication logic
	// Authenticate the admin here

	newUser := models.User{}
	if err := Db.Unscoped().Where("1 = 1").Delete(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "All Users Deleted successfully."})
}
