package toolcontroller

import (
	"b-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAll(ctx *gin.Context) {
	var tools []models.Tool

	models.DB.Find(&tools)
	ctx.JSON(http.StatusOK, gin.H{"tools": tools})
}

func FindById(ctx *gin.Context) {
	var tool models.Tool
	id := ctx.Param("id")

	if err := models.DB.First(&tool, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"tool": tool})

}

func Create(ctx *gin.Context) {
	var tool models.Tool

	if err := ctx.ShouldBindJSON(&tool); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Create(&tool).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to create tool"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Data created successfully", "tool": tool})
	}
}

func Update(ctx *gin.Context) {
	var tool models.Tool
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&tool); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&tool).Where("id = ?", id).Updates(&tool).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update tool"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Updated"})
}

func Delete(ctx *gin.Context) {
	var tool models.Tool
	id := ctx.Param("id")
	// if err := ctx.ShouldBindJSON(&tool); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	if models.DB.Delete(&tool, id).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
}
