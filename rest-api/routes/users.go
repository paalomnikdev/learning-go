package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse request data."})
		return
	}

	err = user.Save()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't save user."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created.", "user": gin.H{"email": user.Email, "id": user.ID}})
}

func Login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
