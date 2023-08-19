package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/Akhilstaar/me-my_encryption/router"
	"github.com/Akhilstaar/me-my_encryption/db"
	"github.com/Akhilstaar/me-my_encryption/config"
    
    "github.com/joho/godotenv"
    "net/http"
    "strconv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        panic("Error loading .env file")
    }
    var DbString = os.Getenv(dbString)

    dbInstance := db.InitDB(config.DbString)
    defer dbInstance.Close()

	utils.Randinit()
	store := cookie.NewStore([]byte(config.CfgAdminPass))
    
    r := gin.Default()
	r.Use(sessions.Sessions("adminsession", store))
	router.PuppyRoute(r, *dbInstance)
	if err := r.Run(config.CfgAddr); err != nil {
		fmt.Println("[Error] " + err.Error())
	}
}
