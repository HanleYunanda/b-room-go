package categorycontroller

import (
	"b-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAll(ctx *gin.Context) {
	var categories []models.Category

	models.DB.Find(&categories)
	ctx.JSON(http.StatusOK, gin.H{"categories": categories})
}

func FindById(ctx *gin.Context) {
	var category models.Category
	id := ctx.Param("id")

	if err := models.DB.First(&category, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"category": category})

}

func Create(ctx *gin.Context) {
	var category models.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&category)
	ctx.JSON(http.StatusOK, gin.H{"message": "Data created successfully", "category": category})
}

func Update(ctx *gin.Context) {
	var category models.Category
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&category).Where("id = ?", id).Updates(&category).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Updated"})
}

func Delete(ctx *gin.Context) {
	var category models.Category
	id := ctx.Param("id")
	// if err := ctx.ShouldBindJSON(&category); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	if models.DB.Delete(&category, id).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
}
