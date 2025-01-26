package database

type Recipe struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}
