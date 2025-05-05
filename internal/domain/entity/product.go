package entity

type Product struct {
	ID       uint `gorm:"primarykey" json:"id" example:"1"`
	Name     string
	ImageURL string
	Cost     float32
}
