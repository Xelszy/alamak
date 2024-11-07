package dto

import (
	"base-gin/domain/dao"
	"base-gin/storage"

	"gorm.io/gorm"
)

type BookRepo struct {
	db       *gorm.DB
	Name     string `json:"name" binding:"required,min=2,max=48"`
	Subtitle string `json:"city" binding:"required,max=32"`
}

func (r *BookRepo) GetList() ([]dao.Book, error) {
	var books []dao.Book
	ctx, cancel := storage.NewDBContext()
	defer cancel()

	tx := r.db.WithContext(ctx).Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return books, nil
}

func (o *BookRepo) ToEntity() dao.Book {
	var item dao.Book
	item.Title = o.Name
	return item
}

type BookResp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city,omitempty"`
}

func (o *BookResp) FromEntity(item *dao.Book) {
	o.ID = int(item.ID)
	o.Name = item.BookPublisher.Name
}

func (r *BookRepo) Update(id uint, updatedData *dao.Book) error {
	ctx, cancel := storage.NewDBContext()
	defer cancel()

	tx := r.db.WithContext(ctx).Model(&dao.Book{}).Where("id = ?", id).Updates(updatedData)
	return tx.Error
}

type BookUpdateReq struct {
	ID          uint   `json:"-"`
	Title       string `json:"title" binding:"required"`
	Subtitle    string `json:"subtitle" binding:"omitempty,max=48"`
	PublisherID uint   `json:"publisher_id" binding:"required"`
}

type BookCreateReq struct {
	Name string `json:"name" binding:"required,min=2,max=48"`
	City string `json:"city" binding:"required,max=32"`
}

func (o *BookCreateReq) ToEntity() dao.Book {
	var item dao.Book
	item.Title = o.Name
	item.Subtitle = &o.City

	return item
}

func (o *BookCreateReq) ToRepo() *BookRepo {
	return &BookRepo{
		Name:     o.Name,
		Subtitle: o.City,
	}
}
