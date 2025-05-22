package models

import (
	"time"
)

type Game struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Title         string     `json:"title" gorm:"size:255;not null;index"`
	Description   string     `json:"description" gorm:"type:text"`
	Developer     string     `json:"developer" gorm:"size:255"`
	Publisher     string     `json:"publisher" gorm:"size:255"`
	ReleaseDate   time.Time  `json:"release_date"`
	Genres        []Genre    `json:"genres" gorm:"many2many:game_genres;"`
	Platforms     []Platform `json:"platforms" gorm:"many2many:game_platforms;"`
	ImageURL      string     `json:"image_url" gorm:"size:255"`
	AverageRating float64    `json:"average_rating" gorm:"type:decimal(3,2)"`
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
	MinRating *float64 `form:"min_rating"`
	SortBy    string   `form:"sort_by"`
	SortOrder string   `form:"sort_order"`
	Page      int      `form:"page" default:"1"`
	PageSize  int      `form:"page_size" default:"10"`
}

type GameResponse struct {
	Games      []Game `json:"games"`
	TotalCount int64  `json:"total_count"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	TotalPages int    `json:"total_pages"`
}
