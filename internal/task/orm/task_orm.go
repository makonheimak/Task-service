package orm

type Task struct {
	ID     int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Task   string `json:"task"`
	UserID int64  `json:"user_id"`
}
