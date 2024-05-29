package wiredemo

import (
	"fmt"
	"net/http"
)

// User 是用户模型
type User struct {
	ID    int
	Name  string
	Email string
}

// UserRepository 是用户存储库
type UserRepository struct {
	DB *Database
}

// NewUserRepository 创建一个新的用户存储库
func NewUserRepository(db *Database) *UserRepository {
	return &UserRepository{DB: db}
}

// GetUserByID 根据ID获取用户
func (repo *UserRepository) GetUserByID(id int) (*User, error) {
	user := &User{}
	err := repo.DB.QueryRow("SELECT id, name, email FROM user WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserService 是用户服务
type UserService struct {
	Repo *UserRepository
}

// NewUserService 创建一个新的用户服务
func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

// GetUserByID 根据ID获取用户
func (service *UserService) GetUserByID(id int) (*User, error) {
	return service.Repo.GetUserByID(id)
}

// UserController 是用户控制器
type UserController struct {
	Service *UserService
}

// NewUserController 创建一个新的用户控制器
func NewUserController(service *UserService) *UserController {
	return &UserController{Service: service}
}

// GetUserByID 处理获取用户的HTTP请求
func (controller *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := 1 // 假设我们从请求中获取了用户ID
	user, err := controller.Service.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User: %v\n", user)
}
