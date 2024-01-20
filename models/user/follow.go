package user

type Follow struct {
	ID       uint `json:"id" gorm:"primaryKey" gorm:"comment:ID"`
	UserID   uint `json:"user_id" gorm:"not null;comment:用户ID"`
	FollowID uint `json:"follow_id" gorm:"not null;comment:关注用户ID"`
}
