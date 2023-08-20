package controllers

import (
	"github.com/Akhilstaar/me-my_encryption/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"gorm.io/gorm"
	"errors"

)

func UserFirstLogin(c *gin.Context) {
	// TODO: Implement user authentication logic
	// Authenticate the user here

	// Validate the input format
	info := new(models.TypeUserFirst)
	if err := c.BindJSON(info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input data format."})
		return
	}

	// See U later ;) ...
	// user := models.User{}
	// record := Db.Model(&user).Select("auth_c").Where("id = ?", info.Id)
	// if record.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error})
	// 	return
	// }

	// Insert the user into the database

	c.JSON(http.StatusCreated, gin.H{"message": "recorded"})
}

func SendHeart(c *gin.Context) {
	// TODO: Implement user authentication logic
	// Authenticate the user here

	info := new(models.SendHeartFirst)
	if err := c.BindJSON(info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input data format."})
		return
	}

	if info.ENC1 != "" && info.SHA1 != "" {
		newheart1 := models.SendHeart{
			SHA:            info.SHA1,
			ENC:            info.ENC1,
			GenderOfSender: info.GenderOfSender,
		}
		if err := Db.Create(&newheart1).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}

	if info.ENC2 != "" && info.SHA2 != "" {
		newheart2 := models.SendHeart{
			SHA:            info.SHA2,
			ENC:            info.ENC2,
			GenderOfSender: info.GenderOfSender,
		}
		if err := Db.Create(&newheart2).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error1": err})
			return
		}
	}

	if info.ENC3 != "" && info.SHA3 != "" {
		newheart3 := models.SendHeart{
			SHA:            info.SHA3,
			ENC:            info.ENC3,
			GenderOfSender: info.GenderOfSender,
		}
		if err := Db.Create(&newheart3).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error2": err})
			return
		}
	}

	if info.ENC4 != "" && info.SHA4 != "" {
		newheart4 := models.SendHeart{
			SHA:            info.SHA4,
			ENC:            info.ENC4,
			GenderOfSender: info.GenderOfSender,
		}
		if err := Db.Create(&newheart4).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error3": err})
			return
		}
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Hearts Sent Successfully !!"})

}

func HeartClaim(c *gin.Context) {
	// TODO: Implement user authentication logic
	// Authenticate the user here

	info := new(models.VerifyHeartClaim)
	if err := c.BindJSON(info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input data format."})
		return
	}

	heartModel := models.SendHeart{}

	verifyheart := Db.Model(&heartModel).Where("sha = ? AND enc = ?", info.SHA, info.Enc).First(&heartModel)
	if verifyheart.Error != nil {
		if errors.Is(verifyheart.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid Heart Claim Request."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": verifyheart.Error.Error()})
		}
		return
	}
	// If the db has record of sha and enc then remove it from the record and add the sha, enc to userId
	if err := Db.Model(&heartModel).Where("sha = ? AND enc = ?", info.SHA, info.Enc).Unscoped().Delete(&heartModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

		// need to change the hardcoded userr string to userId from auth token.
	heartclaim := models.HeartClaims{
		Id: info.Enc,
		SHA: info.SHA,
		Roll:  "userr",
	}
	if err := Db.Create(&heartclaim).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// TODO: Implement "SendClaimedHeartBack" token logic
	// generate a token for "SendClaimedHeartBack" which is valid for 10? mins.

	c.JSON(http.StatusAccepted, gin.H{"message": "Heart Claim Success"})
}

// TODO: Current issue is that if the user changes the enc of the claimed hash(which is very timeconsuming btw ;), there is no way to verify here.
// Why not just add a time window of 10? mins in which the heartback can be accessed.
// So, what are the odds that user gets a heart within 10 mins of submitting its hearts ?.
// Even if the user gets it, what are the odds that user will be able to Intercept the request and make a claim with "enc" which is encoded with pub key of user's 5th choice ?
func SendClaimedHeartBack(c *gin.Context) {
	// TODO: Modify this function to handle multiple concatenated json inputs

	// TODO: Implement user authentication logic
	// Authenticate the user here

	// Authenticate the "SendClaimedHeartBack" token

	info := new(models.UserReturnHearts)
	if err := c.BindJSON(info).Error; err != nil {
		c.JSON(http.StatusMisdirectedRequest, gin.H{"error": "Invalid input data format."})
		return
	}

	returnheart := models.ReturnHearts{
		SHA: info.SHA,
		ENC: info.ENC,
	}

	if err := Db.Create(&returnheart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid tokens SHA/ENC or Server side error. Please try again."})
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Return Heart sent"})
}
