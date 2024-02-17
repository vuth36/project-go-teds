// Trong thư mục internal/transport

package transport

import (
	"net/http"
	"strconv"
	"yourapp/internal/app"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	userUseCase *app.UserUseCase
	renderer    *render.Renderer // Thêm renderer cho templates
}

func NewHTTPHandler(userUC *app.UserUseCase, renderer *render.Renderer) *HTTPHandler {
	return &HTTPHandler{userUseCase: userUC, renderer: renderer}
}

func (h *HTTPHandler) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Xử lý lỗi nếu cần
	}

	user, err := h.userUseCase.GetUserByID(userID)
	if err != nil {
		// Xử lý lỗi nếu cần
	}

	// Render template với dữ liệu người dùng
	h.renderer.HTML(c.Writer, http.StatusOK, "user.html", gin.H{
		"ID":       user.ID,
		"Username": user.Username,
		"Email":    user.Email,
		// Các trường dữ liệu khác của người dùng
	})
}
