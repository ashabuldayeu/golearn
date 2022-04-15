package api

import (
	"github.com/gin-gonic/gin"
	//"githum.com/ashabuldayeu/golearn/go-gin/models"
)

func GetProducts(c *gin.Context) {
	c.JSON(models.strgProducts)
}

func DeleteProduct(c *gin.Context) {

}

func AddProduct(c *gin.Context) {

}
