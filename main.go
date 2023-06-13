package main

import (
	categorycontroller "b-room/controllers/CategoryController"
	reservationcontroller "b-room/controllers/ReservationController"
	roomcontroller "b-room/controllers/RoomController"
	toolcontroller "b-room/controllers/ToolController"
	usercontroller "b-room/controllers/UserController"
	"b-room/models"

	"github.com/gin-gonic/gin"
)

func main() {
	root := gin.Default()
	models.ConnectDB()

	root.GET("api/user", usercontroller.FindAll)
	root.GET("api/user/:id", usercontroller.FindById)
	root.POST("api/user", usercontroller.Create)
	root.PUT("api/user/:id", usercontroller.Update)
	root.DELETE("api/user/:id", usercontroller.Delete)

	root.GET("api/category", categorycontroller.FindAll)
	root.GET("api/category/:id", categorycontroller.FindById)
	root.POST("api/category", categorycontroller.Create)
	root.PUT("api/category/:id", categorycontroller.Update)
	root.DELETE("api/category/:id", categorycontroller.Delete)

	root.GET("api/room", roomcontroller.FindAll)
	root.GET("api/room/:id", roomcontroller.FindById)
	root.POST("api/room", roomcontroller.Create)
	root.PUT("api/room/:id", roomcontroller.Update)
	root.DELETE("api/room/:id", roomcontroller.Delete)

	root.GET("api/reservation", reservationcontroller.FindAll)
	root.GET("api/reservation/:id", reservationcontroller.FindById)
	root.POST("api/reservation", reservationcontroller.Create)
	root.PUT("api/reservation/:id", reservationcontroller.Update)
	root.DELETE("api/reservation/:id", reservationcontroller.Delete)

	root.GET("api/tool", toolcontroller.FindAll)
	root.GET("api/tool/:id", toolcontroller.FindById)
	root.POST("api/tool", toolcontroller.Create)
	root.PUT("api/tool/:id", toolcontroller.Update)
	root.DELETE("api/tool/:id", toolcontroller.Delete)

	root.Run()
}
