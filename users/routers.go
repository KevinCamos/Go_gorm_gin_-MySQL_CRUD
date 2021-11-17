package users

import (
	// "errors"
	// "starbars/common"
	"github.com/gin-gonic/gin"
	"fmt"
	// "net/http"
	// "strconv"
)



func UserRegister(router *gin.RouterGroup) {
	router.POST("/", UserCreate)
	// router.POST("/login", UsersLogin)
}

func UserCreate(c *gin.Context) {
	fmt.Println(c)

	// articleModelValidator := NewArticleModelValidator()
	// if err := articleModelValidator.Bind(c); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
	// 	return
	// }
	// //fmt.Println(articleModelValidator.articleModel.Author.UserModel)

	// if err := SaveOne(&articleModelValidator.articleModel); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	// 	return
	// }
	// serializer := ArticleSerializer{c, articleModelValidator.articleModel}
	// c.JSON(http.StatusCreated, gin.H{"article": serializer.Response()})
}
