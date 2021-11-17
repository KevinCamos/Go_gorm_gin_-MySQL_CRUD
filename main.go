package main

import (
 	// "starbars/Config"
	
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	
	"starbars/common"
	"starbars/Users"
) 

var err error
func Migrate(db *gorm.DB) {
	// users.AutoMigrate()
	// db.AutoMigrate(&articles.ArticleModel{})
	// db.AutoMigrate(&articles.TagModel{})
	// db.AutoMigrate(&articles.FavoriteModel{})
	// db.AutoMigrate(&articles.ArticleUserModel{})
	// db.AutoMigrate(&articles.CommentModel{})
}

func main() {

	fmt.Println("Entra al main:", "linea 22")
	db := common.Init()
	Migrate(db)
	defer db.Close()



	
	r := gin.Default()

	v1 := r.Group("/api")
	users.UserRegister(v1.Group("/users"))










	r.Run()

/* 	r := Routes.SetupRouter()
	//running
	r.Run() */
}
