package controller

import (
	"tick/service"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context){

	op := c.Param("operation")
	
	switch op{
	case "SignUp":
		service.SignUp(c.Request,c.Writer)
	case "SignIn":
		service.SignIn(c.Request,c.Writer)
	case "EditProfile":
		service.EditProfile(c.Request,c.Writer)
	case "ViewProfile":
		service.ViewProfile(c.Request,c.Writer)	
	case "UserList":
		service.UserList(c.Request,c.Writer)
	case "DeleteUser":
		service.DeleteUser(c.Request,c.Writer)				
	}

}