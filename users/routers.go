package users

import (
	// "errors"
	"starbars/common"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	// "strconv"
)



func UserRegister(router *gin.RouterGroup) {
	router.POST("/", UserCreate)
	// router.POST("/login", UsersLogin)
}

func UserCreate(c *gin.Context) {
	fmt.Println(c)
	fmt.Println("eh", "eh")

	userModelValidator := NewUserModelValidator()
	fmt.Println(userModelValidator)

	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

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





/* package users

import (
	// "errors"
  	//"starbars/common"
	  "starbars/config"

	"fmt"
	"github.com/gin-gonic/gin"
	
	_ "github.com/go-sql-driver/mysql"
	// "net/http"
	// "strconv"
)



func UserRegister(router *gin.RouterGroup) {
	router.POST("/", CreateUser)
	// router.POST("/login", UsersLogin)
}



//CreateUser ... Insert New data
func CreateUser(data interface{}) (err error) {
	
	fmt.Println("entra a create User","eh")
	if err = config.DB.Create(data).Error; err != nil {
		return err
	}
	return nil
} */