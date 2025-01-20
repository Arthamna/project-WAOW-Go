package repositories

import (
    "errors"
    "gorm.io/gorm"
    "project_article/internal/models"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
    if err := r.db.Create(user).Error; err != nil {
        return nil, err 
    }
    return user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    result := r.db.Where("email = ?", email).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil 
        }
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) FindAll() ([]models.User, error) {
    var users []models.User
    err := r.db.Find(&users).Error
    return users, err
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
    var user models.User
    result := r.db.First(&user, "user_id = ?", id)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, result.Error
    }
    return &user, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {
    if err := r.db.Save(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

func (r *UserRepository) Delete(id string) error {
    return r.db.Delete(&models.User{}, "id = ?", id).Error
}