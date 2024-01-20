package user

type USubscribe struct {
	ID     uint `json:"id" gorm:"primaryKey" gorm:"comment:ID"`
	TypeID uint `json:"type_id" gorm:"comment:类型ID"`
	UserID uint `json:"user_id" gorm:"comment:用户ID"`
}
