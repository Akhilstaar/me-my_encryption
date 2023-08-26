package controllers

import (
	"net/http"
	"github.com/Akhilstaar/me-my_encryption/models"
	"github.com/gin-gonic/gin"
)

func FetchHearts(c *gin.Context) {
    var heart models.SendHeart
    var hearts []models.FetchHeartsFirst
    // Fetch only required columns from the database
    fetchheart := Db.Model(&heart).Select("enc", "gender_of_sender").Find(&hearts)
    
    if fetchheart.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No hearts to fetch."})
        return
    }

    c.JSON(http.StatusOK, hearts)
}

// func FetchReturnHearts(c *gin.Context) {
// 	heartModel := models.ReturnHearts{}
//     var hearts []models.ReturnHearts
// 	fetchheart := Db.Model(&heartModel).Select("enc").Find(hearts)
// 	if fetchheart.Error != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error" : "No hearts to fetch."})
//         return
// 	}

// 	c.JSON(http.StatusOK, hearts)
// }