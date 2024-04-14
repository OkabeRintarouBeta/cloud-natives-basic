package book

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DTO struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"Author"`
	PublishedDate string `json:"published_date"`
	ImageURL      string `json:"image_url"`
	Description   string `json:"description"`
}

type Form struct {
	Title         string `json:"title" validate:"required,max=255"`
	Author        string `json:"author" validate:"required,alphaspace,max=255"`
	PublishedDate string `json:"published_date" validate:"required,datetime=2006-01-02"`
	ImageURL      string `json:"image_url" validate:"url"`
	Description   string `json:"description"`
}

type Book struct {
	ID            uuid.UUID `gorm:"primarykey"`
	Title         string
	Author        string
	PublishedDate time.Time
	ImageURL      string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Books []*Book
