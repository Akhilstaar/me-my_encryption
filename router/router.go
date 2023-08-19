package router

import (
	"net/http"
	
	// "github.com/Akhilstaar/me-my_encryption/controllers"
	"github.com/Akhilstaar/me-my_encryption/db"
	"github.com/gin-gonic/gin"
)

func PuppyRoute(r *gin.Engine, db db.PuppyDb) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "Hello from the other side!")
	})

	// controllers.Db = db

	// User administration
	// users := r.Group("/users")
	// {
	// 	users.POST("/login/first", controllers.UserFirst)
	// 	users.GET("/mail/:id", controllers.UserMail)
	// }

	// // Session administration
	// session := r.Group("/session")
	// {
	// 	session.POST("/login", controllers.SessionLogin)
	// 	session.GET("/logout", controllers.SessionLogout)
	// }

	// // admin
	// admin := r.Group("/admin")
	// {
	// 	admin.GET("/user/deleteallusers", controllers.DeleteAllUsers)
	// 	admin.POST("/user/new", controllers.UserNew)
	// 	admin.POST("/user/delete", controllers.UserDelete)
	// }

}