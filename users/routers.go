
package users

import (
	// "errors"
	"starbars/config"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	// "strconv"
)



func UserRegister(router *gin.RouterGroup) {
	router.POST("/", UserCreate)
	// router.GET("/", GetAllUSers)
	// router.POST("/login", UsersLogin)
}

func UserCreate(c *gin.Context) {
	var user UserModel
	// var err
	c.BindJSON(&user)
	if err := config.DB.Create(&user).Error; err != nil {
		fmt.Println(err.Error(), "Ací hi ha un error, routers.go de users")
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, user)
}
 

	/* var json UserModel
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(json)

	fmt.Println("eh", "eh")


	if err := config.DB.Create(json).Error; err != nil {
		fmt.Println(err, "err")

		return
	}
	return */
// }
 