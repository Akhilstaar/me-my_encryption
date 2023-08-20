package models

import (
	"gorm.io/gorm"
)

type SendHeartFirst struct {
	GenderOfSender string `json:"genderOfSender" binding:"required"`
	ENC1           string `json:"enc1" binding:"required"`
	SHA1           string `json:"sha1" binding:"required"`
	ENC2           string `json:"enc2" binding:"required"`
	SHA2           string `json:"sha2" binding:"required"`
	ENC3           string `json:"enc3" binding:"required"`
	SHA3           string `json:"sha3" binding:"required"`
	ENC4           string `json:"enc4" binding:"required"`
	SHA4           string `json:"sha4" binding:"required"`
}

type VerifyHeartClaim struct {
	Enc string `json:"enc" binding:"required"`
	SHA string `json:"sha" binding:"required"`
}

type FetchHeartsFirst struct {
	Enc            string `json:"enc"`
	GenderOfSender string `json:"genderOfSender"`
}

// gorm.Model represents the structure of our resource in db
type (
	SendHeart struct {
		gorm.Model
		SHA            string `json:"sha" bson:"sha" gorm:"unique"`
		ENC            string `json:"enc" bson:"enc"`
		GenderOfSender string `json:"genderOfSender" bson:"gender"`
	}
)

type (
	// unique not needed in SHA but still...
	HeartClaims struct {
		gorm.Model
		ENC string `json:"enc" bson:"enc" gorm:"unique"`
		SHA string `json:"sha" bson:"sha" gorm:"unique"`
		Id  string `json:"id"`
	}
)

// --------- Returning Heart Below ---------

type UserReturnHearts struct {
	SHA string `json:"sha" bson:"sha" binding:"required"`
	ENC string `json:"enc" bson:"enc" binding:"required"`
}

type (
	ReturnHearts struct {
		gorm.Model
		SHA string `json:"sha" bson:"sha" gorm:"unique"`
		ENC string `json:"enc" bson:"enc" gorm:"unique"`
	}
)
