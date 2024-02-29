package supplier

import "time"

type Supplier struct {
	ID           int
	KodeSupplier string
	NamaSupplier string
	Alamat       string
	NoHp         string
	Email        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Supplier) TableName() string {
	return "supplier"
}
