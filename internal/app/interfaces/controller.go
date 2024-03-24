package interfaces

import "github.com/gin-gonic/gin"

type IControllerUser interface {
	GetUserByID(c *gin.Context) 
}