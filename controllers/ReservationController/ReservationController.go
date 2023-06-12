package reservationcontroller

import (
	"b-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindAll(ctx *gin.Context) {
	var reservations []models.Reservation

	models.DB.Joins("User").Joins("Room").Joins("Room.Category").Find(&reservations)
	ctx.JSON(http.StatusOK, gin.H{"reservations": reservations})
}

func FindById(ctx *gin.Context) {
	var reservation models.Reservation
	id := ctx.Param("id")

	if err := models.DB.Joins("User").Joins("Room").Joins("Room.Category").First(&reservation, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"reservation": reservation})

}

func Create(ctx *gin.Context) {
	var reservation models.Reservation

	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Create(&reservation).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error creating reservation"})
		return
	} else {
		id := reservation.Id
		models.DB.Joins("User").Joins("Room").Joins("Room.Category").First(&reservation, id)
		ctx.JSON(http.StatusOK, gin.H{"message": "Data created successfully", "reservation": reservation})
	}
}

func Update(ctx *gin.Context) {
	var reservation models.Reservation
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&reservation).Where("id = ?", id).Updates(&reservation).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Error updating reservation"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Updated"})
}

func Delete(ctx *gin.Context) {
	var reservation models.Reservation
	id := ctx.Param("id")
	// if err := ctx.ShouldBindJSON(&reservation); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	if models.DB.Delete(&reservation, id).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
}
