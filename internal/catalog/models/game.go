package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"size:255;not null;index"`
	Description string         `json:"description" gorm:"type:text"`
	Developer   string         `json:"developer" gorm:"size:255"`
	Publisher   string         `json:"publisher" gorm:"size:255"`
	ReleaseDate time.Time      `json:"release_date"`
	Genres      []Genre        `json:"genres" gorm:"many2many:game_genres;"`
	Platforms   []Platform     `json:"platforms" gorm:"many2many:game_platforms;"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2)"`
	ImageURL    string         `json:"image_url" gorm:"size:255"`
	AverageRating float64      `json:"average_rating" gorm:"type:decimal(3,2)"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type Genre struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:100;not null;uniqueIndex"`
}

type Platform struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:100;not null;uniqueIndex"`
}

type GameFilter struct {
	Title     string   `form:"title"`
	Developer string   `form:"developer"`
	Publisher string   `form:"publisher"`
	Genres    []string `form:"genres"`
	Platforms []string `form:"platforms"`
	MinPrice  *float64 `form:"min_price"`
	MaxPrice  *float64 `form:"max_price"`
	MinRating *float64 `form:"min_rating"`
	SortBy    string   `form:"sort_by"`
	SortOrder string   `form:"sort_order"`
	Page      int      `form:"page"`
	PageSize  int      `form:"page_size"`
}

type GameResponse struct {
	Games      []Game `json:"games"`
	TotalCount int64  `json:"total_count"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	TotalPages int    `json:"total_pages"`
}
