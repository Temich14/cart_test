package entity

type Product struct {
	ID       string `gorm:"primary_key"`
	Name     string
	ImageURL string
	Cost     float32
}
