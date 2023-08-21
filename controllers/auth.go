package controllers

import (
	"fmt"
	"net/http"
	"os"
	"log"
	// "os"
	// "time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthenticateAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authCookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
				c.Abort()
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSigningKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"].(string)
			if userID!=os.Getenv("AdminId"){
				unauthorizedMessage := fmt.Sprintf("Unauthorized Login attempt by %s, it will be recorded.", userID)
				log.Println(unauthorizedMessage)
				c.JSON(http.StatusUnauthorized, gin.H{"error": unauthorizedMessage})
				c.Abort()				
			}
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
		}
	}
}

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization cookie missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authCookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
				c.Abort()
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSigningKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"].(string)
			if userID==os.Getenv("AdminId"){
				c.JSON(http.StatusUnauthorized, gin.H{"error": "You Serious ?? , login using Postman or CLI, the frontend is only for normal users."})
				c.Abort()				
			}
			c.Set("user_id", userID)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
		}
	}
}

func AuthenticateUserHeartclaim() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Cookie("HeartBack")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Late HeartClaim cookie missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(authCookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
				c.Abort()
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSigningKey, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			HeartBackUserID := claims["user_id"]
			useridfromJWT,_ := c.Get("user_id")

			if(useridfromJWT.(string) != HeartBackUserID){
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Forged Claim Token"})
				c.Abort()
				return
			}
			c.Set("HeartBackUserID", HeartBackUserID)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
		}
	}
}

