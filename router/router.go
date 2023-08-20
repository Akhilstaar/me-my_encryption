package router

import (
	"net/http"
	
	"github.com/Akhilstaar/me-my_encryption/controllers"
	"github.com/Akhilstaar/me-my_encryption/db"
	"github.com/gin-gonic/gin"
)

func PuppyRoute(r *gin.Engine, db db.PuppyDb) {
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "Hello from the other side!")
	})
	
	// assigning here cuz I'm only inporting controllers here, & considering their size better import them here.
	controllers.Db = db

	// User administration
	users := r.Group("/users")
	{
		users.Use(controllers.AuthenticateUser())
		// users.POST("/login/first", controllers.UserFirstLogin)
		users.POST("/sendheart", controllers.SendHeart)
		users.POST("/claimheart", controllers.HeartClaim)
		users.POST("/sendclaimedheartback", controllers.SendClaimedHeartBack)
		// users.GET("/mail/:id", controllers.UserMail)
	}
	
	// Session administration
	session := r.Group("/session")
	{
		session.POST("/admin/login", controllers.AdminLogin)
		session.POST("/login", controllers.UserLogin)
		session.GET("/logout", controllers.UserLogout)
	}

	// admin
	admin := r.Group("/admin")
	{
		admin.Use(controllers.AuthenticateAdmin())
		admin.GET("/user/deleteallusers", controllers.DeleteAllUsers)
		admin.POST("/user/new", controllers.AddNewUser)
		admin.POST("/user/delete", controllers.DeleteUser)
	}

}