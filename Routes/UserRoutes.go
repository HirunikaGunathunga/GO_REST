package routes

import (
	"fmt"
	"net/http"

	model "github.com/HirunikaGunathunga/GO_REST/Model"
	utils "github.com/HirunikaGunathunga/GO_REST/Utils"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var u model.Users
	err := context.ShouldBindJSON(&u) //Bind data from request
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the action"})
		return
	}

	err = u.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created", "User: ": u})
}

func Login(context *gin.Context) {
	var u model.Users
	err := context.ShouldBindJSON(&u) //Bind data from request
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the action"})
		return
	}
	err = u.ConfirmUser()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.JwtToken(u.Email, u.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not auth the user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login Successful", "token": token})
}
