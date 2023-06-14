package main

import (
	authcontroller "b-room/controllers/AuthController"
	categorycontroller "b-room/controllers/CategoryController"
	reservationcontroller "b-room/controllers/ReservationController"
	roomcontroller "b-room/controllers/RoomController"
	toolcontroller "b-room/controllers/ToolController"
	usercontroller "b-room/controllers/UserController"
	"b-room/middleware"
	"b-room/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	root := gin.Default()
	models.ConnectDB()

	root.POST("api/login", authcontroller.Login)
	root.POST("api/logout", authcontroller.Logout)

	root.POST("api/register", usercontroller.Create)

	root.GET("api/user", middleware.RequireAuth, usercontroller.FindAll)
	root.GET("api/user/:id", middleware.RequireAuth, usercontroller.FindById)
	// root.POST("api/user", usercontroller.Create)
	root.PUT("api/user/:id", middleware.RequireAuth, usercontroller.Update)
	root.DELETE("api/user/:id", middleware.RequireAuth, usercontroller.Delete)

	root.GET("api/category", middleware.RequireAuth, categorycontroller.FindAll)
	root.GET("api/category/:id", middleware.RequireAuth, categorycontroller.FindById)
	root.POST("api/category", middleware.RequireAuth, categorycontroller.Create)
	root.PUT("api/category/:id", middleware.RequireAuth, categorycontroller.Update)
	root.DELETE("api/category/:id", middleware.RequireAuth, categorycontroller.Delete)

	root.GET("api/room", middleware.RequireAuth, roomcontroller.FindAll)
	root.GET("api/room/:id", middleware.RequireAuth, roomcontroller.FindById)
	root.POST("api/room", middleware.RequireAuth, roomcontroller.Create)
	root.PUT("api/room/:id", middleware.RequireAuth, roomcontroller.Update)
	root.DELETE("api/room/:id", middleware.RequireAuth, roomcontroller.Delete)

	root.GET("api/reservation", middleware.RequireAuth, reservationcontroller.FindAll)
	root.GET("api/reservation/:id", middleware.RequireAuth, reservationcontroller.FindById)
	root.POST("api/reservation", middleware.RequireAuth, reservationcontroller.Create)
	root.PUT("api/reservation/:id", middleware.RequireAuth, reservationcontroller.Update)
	root.DELETE("api/reservation/:id", middleware.RequireAuth, reservationcontroller.Delete)

	root.GET("api/tool", middleware.RequireAuth, toolcontroller.FindAll)
	root.GET("api/tool/:id", middleware.RequireAuth, toolcontroller.FindById)
	root.POST("api/tool", middleware.RequireAuth, toolcontroller.Create)
	root.PUT("api/tool/:id", middleware.RequireAuth, toolcontroller.Update)
	root.DELETE("api/tool/:id", middleware.RequireAuth, toolcontroller.Delete)

	root.Run()
}
