// // Trong thư mục internal/storage

package storage

// type UserRepository struct {
// 	db *YourDatabase // Giao tiếp với cơ sở dữ liệu
// }

func NewUserRepository(db *YourDatabase) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	// Lấy thông tin người dùng từ cơ sở dữ liệu
	// Sử dụng SQL hoặc ORM
}

// // Các phương thức khác để tương tác với người dùng trong cơ sở dữ liệu
