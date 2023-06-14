package roomcontroller

import (
	"b-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAll(ctx *gin.Context) {
	var rooms []models.Room

	models.DB.Joins("Category").Find(&rooms)
	ctx.JSON(http.StatusOK, gin.H{"rooms": rooms})
}

func FindById(ctx *gin.Context) {
	var room models.Room
	id := ctx.Param("id")

	if err := models.DB.Joins("Category").First(&room, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"room": room})

}

func Create(ctx *gin.Context) {
	var room models.Room

	if err := ctx.ShouldBindJSON(&room); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Create(&room).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to create room"})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Data created successfully", "room": room})
	}
}

func Update(ctx *gin.Context) {
	var room models.Room
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&room); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&room).Where("id = ?", id).Updates(&room).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Failed to update room"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Updated"})
}

func Delete(ctx *gin.Context) {
	var room models.Room
	id := ctx.Param("id")
	// if err := ctx.ShouldBindJSON(&room); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	if models.DB.Delete(&room, id).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
}
