package supplier

import "gorm.io/gorm"

type Repository interface {
	Save(supplier Supplier) (Supplier, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(supplier Supplier) (Supplier, error) {
	err := r.db.Create(&supplier).Error
	if err != nil {
		return supplier, err
	}

	return supplier, nil

}
