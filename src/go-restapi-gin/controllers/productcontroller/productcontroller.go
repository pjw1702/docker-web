package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pjw1702/go-restapi-gin/models"
	"gorm.io/gorm"
)

// GET
func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"products": products})
}

// GET
func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id") // main()의 /api/pjw/product/:id에서, DB의 id 값이 파라미터로 적용된다는 뜻 (ex: localhost:8080/api/pjw/product/2)

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data is not valuable"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

// POST
func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product})
}

// PUT
func Update(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not found Id of product you want to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully update data of the product"})
}

// DELETE
func Delete(c *gin.Context) {
	var product models.Product

	var input struct {
		Id json.Number
	}

	// input := map[string]string{"id": "0"}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// id, _ := strconv.ParseInt(input["id"], 10, 64)
	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Not found Id of product you want to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete data of the product"})
}
