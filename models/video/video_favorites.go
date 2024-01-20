package video

type VFavoritesVideo struct {
	ID          uint `json:"id" gorm:"primaryKey" gorm:"comment:ID"`
	FavoritesID uint `json:"favorites_id" gorm:"not null;index;comment:收藏夹ID"`
	VideoID     uint `json:"video_id" gorm:"comment:视频ID"`
	UserID      uint `json:"user_id" gorm:"not null;index;comment:用户ID"`
}
