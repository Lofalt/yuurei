package controller

import "github.com/gin-gonic/gin"

type RESTfulcontroller interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Show(c *gin.Context)
	Delete(c *gin.Context)
}
