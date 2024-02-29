package refid

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(refid RefId) (RefId, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(refid RefId) (RefId, error) {
	err := r.db.Create(&refid).Error
	if err != nil {
		return refid, err
	}
	return refid, nil
}
