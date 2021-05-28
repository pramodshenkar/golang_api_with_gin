package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/golang_todo_api/forms"
	"github.com/pramodshenkar/golang_todo_api/helpers"
	"github.com/pramodshenkar/golang_todo_api/models"
	"github.com/pramodshenkar/golang_todo_api/services"
)

var userModel = new(models.UserModel)

type UserController struct{}

func (u *UserController) Signup(c *gin.Context) {
	var data forms.SignupUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}

	result, _ := userModel.GetUserByEmail(data.Email)

	if result.Email != "" {
		c.JSON(403, gin.H{"message": "Email is already in use"})
		c.Abort()
		return
	}

	err := userModel.Signup(data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}
	c.JSON(201, gin.H{"message": "New user account registered"})
}

func (u *UserController) Login(c *gin.Context) {
	var data forms.LoginUserCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	result, err := userModel.GetUserByEmail(data.Email)

	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem logging into your account"})
		c.Abort()
		return
	}

	hashedPassword := []byte(result.Password)
	password := []byte(data.Password)

	err = helpers.PasswordCompare(password, hashedPassword)

	if err != nil {
		c.JSON(403, gin.H{"message": "Invalid user credentials"})
		c.Abort()
		return
	}

	jwtToken, err2 := services.GenerateToken(data.Email)

	// If we fail to generate token for access
	if err2 != nil {
		c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Log in success", "token": jwtToken})
}
