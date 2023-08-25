package controllers

import (
	"net/http"
	"github.com/Akhilstaar/me-my_encryption/models"
	"github.com/gin-gonic/gin"
)

func FetchHearts(c *gin.Context) {
    var hearts []models.SendHeart
    fetchheart := Db.Model(&hearts).Where("1 = 1").Find(&hearts)
    
    if fetchheart.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No hearts to fetch."})
        return
    }
    
    var simplifiedHearts []map[string]interface{}
    for _, heart := range hearts {
        simplifiedHeart := map[string]interface{}{
            "enc":           heart.ENC,
            "genderOfSender": heart.GenderOfSender,
        }
        simplifiedHearts = append(simplifiedHearts, simplifiedHeart)
    }
    
    c.JSON(http.StatusOK, simplifiedHearts)
}


func FetchReturnHearts(c *gin.Context) {
	// TODO: Implement user authentication logic
	// Authenticate the user here


	heartModel := models.SendHeart{}
	fetchheart := Db.Model(&heartModel).Select("enc","genderOfSender")
	if fetchheart.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "No hearts to fetch."})
	}

	c.JSON(http.StatusAccepted, gin.H{"hearts" : fetchheart})
}