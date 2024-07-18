package model

type VideoInMongo struct {
	ID        int64  `bson:"id"`
	UserID    int64  `bson:"user_id"`
	Title     string `bson:"title"`
	PlayURL   string `bson:"play_url"`
	CoverURL  string `bson:"cover_url"`
	Label     string `bson:"label"`
	Category  string `bson:"category"`
	Timestamp int64  `bson:"timestamp"`
}
