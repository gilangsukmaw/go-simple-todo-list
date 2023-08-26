package repositries

import (
	"gorm.io/gorm"
	"simple-todo-list/internal/consts"
	"simple-todo-list/internal/entities"
)

// Repository interface allows us to access the CRUD Operations in sql here.
type UserRepository interface {
	CreateUser(user *entities.User) error
	FindOne(username string) (*entities.User, error)
}
type userRepository struct {
	db *gorm.DB
}

// NewRepo is the single instance repo that is being created.
func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// CreateUser is a gorm repository that helps to create User
func (r *userRepository) CreateUser(user *entities.User) error {
	err := r.db.Create(user).Error

	return err
}

// ReadBook is a gorm repository that helps to fetch books
func (r *userRepository) FindOne(username string) (*entities.User, error) {
	user := entities.User{}
	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil && err.Error() == consts.SqlNoRow {
		return nil, nil
	}

	return &user, err
}
