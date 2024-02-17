// Trong thư mục cmd

package main

import (
	"yourapp/internal/app"
	"yourapp/internal/storage"
	"yourapp/internal/transport"

	"github.com/gin-gonic/contrib/render" // Import package render
	"github.com/gin-gonic/gin"
)

func main() {
    // Khởi tạo các thành phần
    db := YourDatabaseConnect() // Kết nối đến cơ sở dữ liệu của bạn
    userRepo := storage.NewUserRepository(db)
    userUC := app.NewUserUseCase(userRepo)

    // Khởi tạo renderer
    renderer := render.New(render.Options{
        Directory:  "templates",
        Extensions: []string{".html"},
    })

    httpHandler := transport.NewHTTPHandler(userUC, renderer)

    // Khởi tạo router Gin
    r := gin.Default()

    // Định nghĩa các routes
    r.GET("/user/:id", httpHandler.GetUserByID)

    // Khởi động server
    r.Run(":8080")
}
