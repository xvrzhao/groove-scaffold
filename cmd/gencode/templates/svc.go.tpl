package handler

import (
	"fmt"
	"{{.module}}/db"
	"{{.module}}/db/model"

	"gorm.io/gorm/clause"
)

type {{.lowerCamelModel}}Svc struct{}

func ({{.lowerCamelModel}}Svc) Page(offset, limit int) ({{.lowerCamelModel}}s []model.{{.upperCamelModel}}, total int64, err error) {
	if err := db.Client.Model(&model.{{.upperCamelModel}}{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	{{.lowerCamelModel}}s = make([]model.{{.upperCamelModel}}, 0, limit)
	if err := db.Client.Limit(limit).Offset(offset).Preload(clause.Associations).Find(&{{.lowerCamelModel}}s).Error; err != nil {
		return nil, 0, err
	}
	return {{.lowerCamelModel}}s, total, nil
}

func ({{.lowerCamelModel}}Svc) Create({{.lowerCamelModel}} *model.{{.upperCamelModel}}) error {
	if err := db.Client.Create({{.lowerCamelModel}}).Error; err != nil {
		return fmt.Errorf("failed create {{.lowerCamelModel}}: %w", err)
	}
	if err := db.Client.Preload(clause.Associations).First({{.lowerCamelModel}}).Error; err != nil {
		return fmt.Errorf("failed to load {{.lowerCamelModel}}: %w", err)
	}
	return nil
}

func ({{.lowerCamelModel}}Svc) Update(id int, {{.lowerCamelModel}} *model.{{.upperCamelModel}}) error {
	if err := db.Client.Where("id = ?", id).Updates({{.lowerCamelModel}}).Error; err != nil {
		return fmt.Errorf("failed create {{.lowerCamelModel}}: %w", err)
	}
	{{.lowerCamelModel}}.ID = uint(id)
	if err := db.Client.Preload(clause.Associations).First({{.lowerCamelModel}}).Error; err != nil {
		return fmt.Errorf("failed to load {{.lowerCamelModel}}: %w", err)
	}
	return nil
}

func ({{.lowerCamelModel}}Svc) Delete(id int) error {
	return db.Client.Delete(&model.{{.upperCamelModel}}{}, id).Error
}
