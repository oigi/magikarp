package model

import (
    "github.com/oigi/Magikarp/config/model"
    "time"
)

type Favorite struct {
    model.Model
    UserID  uint `gorm:"column:user_id;not null" json:"user_id"`
    VideoID uint `gorm:"column:video_id;not null" json:"video_id"`
    Status  int
}

type FavoriteCount struct {
    VideoId       int64
    FavoriteCount int64
    CreatedAt     time.Time
    UpdatedAt     time.Time
}
