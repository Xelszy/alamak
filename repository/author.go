package repository

import (
	"base-gin/domain/dao"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) Create(author *dao.Author) error {
	return r.db.Create(author).Error
}

func (r *AuthorRepository) GetAll() ([]dao.Author, error) {
	var authors []dao.Author
	err := r.db.Find(&authors).Error
	return authors, err
}

func (r *AuthorRepository) GetByID(id uint) (dao.Author, error) {
	var author dao.Author
	err := r.db.First(&author, id).Error
	return author, err
}

func (r *AuthorRepository) Update(author *dao.Author) error {
	return r.db.Save(author).Error
}

func (r *AuthorRepository) Delete(id uint) error {
	return r.db.Delete(&dao.Author{}, id).Error
}
