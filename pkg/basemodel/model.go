package basemodel

import "gorm.io/gorm"

type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt Time           `json:"createdAt"`
	UpdatedAt Time           `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
