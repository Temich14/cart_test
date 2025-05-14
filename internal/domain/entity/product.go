package entity

type Product struct {
	ID       uint    `gorm:"primarykey" json:"id" example:"1"`
	Name     string  `json:"name" example:"iphone 16"`
	ImageURL string  `json:"image_url" example:"https://example.com/example.png"`
	Cost     float32 `json:"cost" example:"799.00"`
}
