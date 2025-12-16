package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `json:"name" validate:"required"`
	Age  int32  `json:"age" validate:"required,min=0,max=120"`
}

func main() {
	r := gin.Default()
	r.POST("/user", func(c *gin.Context) {

		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, "Erro ao converter os dados")
			return
		}

		validate := validator.New()
		err := validate.Struct(user)

		if err != nil {
			c.JSON(http.StatusBadRequest, "Erro ao validar campo")
			return
		}

		c.JSON(http.StatusOK, user)
	})
	r.Run()
}
