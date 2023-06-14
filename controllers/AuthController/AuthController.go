package authcontroller

import (
	"b-room/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var user models.User
	var body struct {
		Email    string
		Password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}

	models.DB.First(&user, "email = ?", body.Email)
	if user.Id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create token",
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenStr, 3600, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"user":    user,
		"token":   tokenStr,
	})
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "Unauthorized", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
