package usercontroller

import (
	"b-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func FindAll(ctx *gin.Context) {
	var users []models.User

	models.DB.Find(&users)
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func FindById(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")

	if err := models.DB.First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})

}

func Create(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(pass)

	models.DB.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func Update(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(pass)

	if models.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Updated"})
}

func Delete(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")
	// if err := ctx.ShouldBindJSON(&user); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }

	if models.DB.Delete(&user, id).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No Data Found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data Deleted"})
}
