package supplier

import "gorm.io/gorm"

type Repository interface {
	Save(supplier Supplier) (Supplier, error)
	FindAll() ([]Supplier, error)
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

func (r *repository) FindAll() ([]Supplier, error) {
	var suppliers []Supplier

	err := r.db.Find(&suppliers).Error

	if err != nil {
		return suppliers, err
	}

	return suppliers, nil

}
