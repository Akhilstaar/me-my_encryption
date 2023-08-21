package controllers

import (
	"net/http"
	"os"
	// "github.com/Akhilstaar/me-my_encryption/models"
	"errors"
	"time"

	"github.com/Akhilstaar/me-my_encryption/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var jwtSigningKey = []byte(os.Getenv("UserjwtSigningKey"))
var heartjwtSigningKey = []byte(os.Getenv("HeartjwtSigningKey"))

func UserLogin(c *gin.Context) {
	info := new(models.UserLogin)
	if err := c.BindJSON(info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input data format."})
		return
	}

	loginmodel := models.User{}
	verifyuser := Db.Model(&loginmodel).Where("id = ? AND pass = ?", info.Id, info.Pass).First(&loginmodel)
	if verifyuser.Error != nil {
		if errors.Is(verifyuser.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid Login Request."})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": verifyuser.Error.Error()})
			return
		}
	}

	token, err := generateJWTToken(info.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT token"})
		return
	}
	expirationTime := time.Now().Add(time.Hour * 24)
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Expires:  expirationTime,
		Path:     "/",
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   false, // Set this to true if you're using HTTPS, false for HTTP
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully !!"})
}

func UserLogout(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Expire the cookie immediately
		Path:     "/",
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged out successfully.",
	})
}

func generateJWTToken(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func generateJWTTokenForHeartBack(userID string) (string, error) {

	claims := jwt.MapClaims{
		"verified": "Absolutely",
		"user_id":  userID,
		"exp":      time.Now().Add(time.Minute * 11).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(heartjwtSigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
