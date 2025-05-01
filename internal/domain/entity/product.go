package entity

type Product struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	ImageURL string
	Cost     float32
}
